package http_controller

import (
	"fmt"
	cmdex "in_mem_storage/internal/application/service/crud_cmd_executor"
	"in_mem_storage/internal/application/service/logger"
	rlim "in_mem_storage/internal/application/service/rate_limiter"
	reqsrv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	req "in_mem_storage/internal/domain/incoming_request/abstraction"
	reqhdnl "in_mem_storage/internal/domain/incoming_request/value_object"
	log "in_mem_storage/internal/domain/log/value_object"
	logrec "in_mem_storage/internal/domain/log/value_object/log_record"
	stfrup "in_mem_storage/internal/domain/log/value_object/stack_frames_up"
	reqlimprod "in_mem_storage/internal/domain/rate_limiter/abstraction"
	rlimobj "in_mem_storage/internal/domain/rate_limiter/value_object"
	commands "in_mem_storage/internal/domain/transaction/command/value_object"
	"time"
)

const (
	WriterErrEvent          log.Event = "[RequestHandlerWriterError]"
	RateLimProducerErrEvent log.Event = "[RateLimitProducerError]"
	RateLimServiceErrEvent  log.Event = "[RateLimitServiceError]"
	CrudCmdProducerErrEvent log.Event = "[CrudCmdProducerError]"
	CmdExecRes              log.Event = "[CrudCommandResult]"
)

func makeLog(
	lvl log.LogLvl,
	event log.Event,
	data string,
	stFramesUp stfrup.StackFramesUp,
) logrec.DefaultLogRecord {
	return logrec.New(lvl, event, "", data, stFramesUp)
}

func logAndWriteError[S ~string](
	logger logger.Logger,
	writer req.Writer[S],
	lvl log.LogLvl,
	event log.Event,
	err error,
) {
	if err != nil {
		logger.Log(makeLog(lvl, event, err.Error(), stfrup.InOuterFn))
		err = writer.Write(S(err.Error()))

		if err != nil {
			logger.Log(makeLog(lvl, event, err.Error(), stfrup.InOuterFn))
		}
	}
}

type ReturningFunc[R, W any] func(
	reqSrv reqsrv.RequestService[R, W],
	cmdEx cmdex.CrudCommandService,
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
	logger logger.Logger,
)

func RateLimiterRoute[
	R reqlimprod.RateLimitProducer,
	S ~string,
	W req.Writer[S],
](path string) ReturningFunc[R, W] {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService,
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
		logger logger.Logger,
	) {
		handler := func(r R, w W) {
			limit, err := r.ProduceRateLim()
			if err != nil {
				logAndWriteError[S](logger, w, log.Error, RateLimProducerErrEvent, err)
				return
			}

			err = rLim.Set(limit.For, limit)
			if err != nil {
				logAndWriteError[S](logger, w, log.Error, RateLimServiceErrEvent, err)
				return
			}

			res := fmt.Sprintf("[RateLimitOperationSuccess] Your rate limit is, %v", limit)
			err = w.Write(S(res))

			logger.Log(makeLog(log.Error, WriterErrEvent, err.Error(), stfrup.Here))
		}
		reqSrv.HandleReqWith(reqhdnl.ReqHandler[R, W]{Path: path, Handle: handler})
	}
	return f
}

func CrudCommandsRoute[
	B any,
	R CrudReq[B],
	S ~string,
	W req.Writer[S],
](path string) ReturningFunc[R, W] {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService,
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
		logger logger.Logger,
	) {
		handle := func(r R, w W) {
			user, now := r.From(), time.Now()

			rateLimit, err := rLim.Get(user)
			if err != nil {
				logAndWriteError[S](logger, w, log.Error, RateLimServiceErrEvent, err)
				defaultZeroLim := rlimobj.RateLimit{For: user, Limit: time.Nanosecond * 0}
				rateLimit = defaultZeroLim
			}

			lastUsed, expiresAfter := rateLimit.LastUsed, rateLimit.Limit
			elapsed := lastUsed.Add(expiresAfter).Sub(now)

			time.Sleep(elapsed)

			cmd, err := r.ProduceCmd()
			if err != nil {
				logAndWriteError[S](logger, w, log.Error, CrudCmdProducerErrEvent, err)
				return
			}

			res := cmdEx.Execute(cmd)
			err = w.Write(S(res.String()))

			logger.Log(makeLog(log.Error, WriterErrEvent, err.Error(), stfrup.Here))
		}
		reqSrv.HandleReqWith(reqhdnl.ReqHandler[R, W]{Path: path, Handle: handle})
	}
	return f
}

func TimeToLiveRoute[R, W any](sleep time.Duration) ReturningFunc[R, W] {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService,
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
		logger logger.Logger,
	) {
		for {
			time.Sleep(sleep)

			now := time.Now()
			expiredRec, err := ttl.Get(now)
			if err != nil {
				continue
			}

			cmd := commands.DeleteCommand{Key: expiredRec.Record}
			res := cmdEx.Execute(cmd)

			logger.Log(makeLog(log.Info, CmdExecRes, res.String(), stfrup.Here))
		}
	}
	return f
}

package crud_controller

import (
	"fmt"
	cmdex "in_mem_storage/application/service/crud_cmd_executor"
	"in_mem_storage/application/service/logger"
	rlim "in_mem_storage/application/service/rate_limiter"
	reqsrv "in_mem_storage/application/service/server"
	ttlserv "in_mem_storage/application/service/time_to_live"
	req "in_mem_storage/domain/incoming_request/abstraction"
	reqhdnl "in_mem_storage/domain/incoming_request/value_object"
	logrec "in_mem_storage/domain/log/value_object/log_record"
	reqlimprod "in_mem_storage/domain/rate_limiter/abstraction"
	rlimobj "in_mem_storage/domain/rate_limiter/value_object"
	"time"
)

const (
	WriterErrEvent          = "[RequestHandlerWriterError]"
	RateLimProducerErrEvent = "[RateLimitProducerError]"
	RateLimServiceErrEvent  = "[RateLimitServiceError]"
	CrudCmdProducerErrEvent = "[CrudCmdProducerError]"
)

func makeLog(event string, err error) logrec.DefaultLogRecord {
	return logrec.New("Error", event, "", err)
}

func logAndWriteError[S ~string](
	logger logger.Logger,
	writer req.Writer[S],
	event string,
	err error,
) {
	logger.Log(makeLog(event, err))

	err = writer.Write(S(err.Error()))

	logger.Log(makeLog(WriterErrEvent, err))
}

type RateLimFunc[
	R reqlimprod.RateLimitProducer,
	S ~string,
	W req.Writer[S],
] func(
	reqSrv reqsrv.RequestService[R, W],
	cmdEx cmdex.CrudCommandService,
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
	logger logger.Logger,
)

func NewRateLimiterRoute[
	R reqlimprod.RateLimitProducer,
	S ~string,
	W req.Writer[S],
](path string) RateLimFunc[R, S, W] {
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
				logAndWriteError[S](logger, w, RateLimProducerErrEvent, err)
				return
			}

			err = rLim.Set(limit.For, limit)
			if err != nil {
				logAndWriteError[S](logger, w, RateLimServiceErrEvent, err)
				return
			}

			res := fmt.Sprintf("[RateLimitOperationSuccess] Your rate limit is, %v", limit)
			err = w.Write(S(res))

			logger.Log(makeLog(WriterErrEvent, err))
		}
		reqSrv.Handle(reqhdnl.ReqHandler[R, W]{Path: path, Handle: handler})
	}
	return f
}

type CrudCommandFunc[
	B any,
	R CrudReq[B],
	S ~string,
	W req.Writer[S],
] func(
	reqSrv reqsrv.RequestService[R, W],
	cmdEx cmdex.CrudCommandService,
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
	logger logger.Logger,
)

func CrudCommandsRoute[
	B any,
	R CrudReq[B],
	S ~string,
	W req.Writer[S],
](path string) CrudCommandFunc[B, R, S, W] {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService,
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
		logger logger.Logger,
	) {
		handle := func(r R, w W) {
			user, now := r.From(), r.Date()

			rateLimit, err := rLim.Get(user)
			if err != nil {
				logAndWriteError[S](logger, w, RateLimServiceErrEvent, err)

				defaultZeroLim := rlimobj.RateLimit{For: user, Limit: time.Nanosecond * 0}
				rateLimit = defaultZeroLim
			}

			lastUsed, expiresAfter := rateLimit.LastUsed, rateLimit.Limit
			elapsed := lastUsed.Add(expiresAfter).Sub(now)

			time.Sleep(elapsed)

			cmd, err := r.ProduceCmd()
			if err != nil {
				logAndWriteError[S](logger, w, CrudCmdProducerErrEvent, err)
				return
			}

			res := cmdEx.Execute(cmd)
			err = w.Write(S(res.String()))

			logger.Log(makeLog(WriterErrEvent, err))
		}
		reqSrv.Handle(reqhdnl.ReqHandler[R, W]{Path: path, Handle: handle})
	}
	return f
}

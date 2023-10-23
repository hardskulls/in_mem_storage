package crud_controller

import (
	cmdex "in_mem_storage/application/service/command_executor"
	rlim "in_mem_storage/application/service/rate_limiter"
	reqsrv "in_mem_storage/application/service/request_handler"
	ttlserv "in_mem_storage/application/service/time_to_live"
	req "in_mem_storage/domain/incoming_request/abstraction"
	reqhdnl "in_mem_storage/domain/incoming_request/value_object"
	reqlimprod "in_mem_storage/domain/rate_limiter/abstraction"
	"in_mem_storage/domain/transaction/command/crud"
	"time"
)

type RateLimFunc[
	C crud.CrudCommand,
	R reqlimprod.RateLimitProducer,
	S ~string,
	W req.Writer[S],
] func(
	reqSrv reqsrv.RequestService[R, W],
	cmdEx cmdex.CrudCommandService[C],
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
)

func NewRateLimiter[
	C crud.CrudCommand,
	R reqlimprod.RateLimitProducer,
	S ~string,
	W req.Writer[S],
](path string) RateLimFunc[C, R, S, W] {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService[C],
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
	) {
		handler := func(r R, w W) {
			limit, err := r.Produce()
			if err != nil {
				_ = w.Write(S(err.Error()))
				return
			}

			err = rLim.Set(limit.For, limit)
			if err != nil {
				_ = w.Write(S(err.Error()))
				return
			}
		}
		reqSrv.Handle(reqhdnl.ReqHandler[R, W]{Path: path, Handle: handler})
	}
	return f
}

type CrudCommandFunc[
	B any,
	C crud.CrudCommand,
	R CrudReq[B, C],
	S ~string,
	W req.Writer[S],
] func(
	reqSrv reqsrv.RequestService[R, W],
	cmdEx cmdex.CrudCommandService[C],
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
)

func DefaultCfg[
	B any,
	C crud.CrudCommand,
	R CrudReq[B, C],
	S ~string,
	W req.Writer[S],
](path string) CrudCommandFunc[B, C, R, S, W] {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService[C],
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
	) {
		handle := func(r R, w W) {
			user, now := r.From(), r.Date()

			rateLimit, err := rLim.Get(user)
			if err != nil {
				_ = w.Write(S(err.Error()))
				return
			}

			lastUsed, expiresAfter := rateLimit.LastUsed, rateLimit.Limit
			elapsed := lastUsed.Add(expiresAfter).Sub(now)

			time.Sleep(elapsed)

			cmd, err := r.ProduceCmd()
			if err != nil {
				_ = w.Write(S(err.Error()))
				return
			}

			res := cmdEx.Execute(cmd)
			_ = w.Write(S(res.String()))
		}
		reqSrv.Handle(reqhdnl.ReqHandler[R, W]{Path: path, Handle: handle})
	}
	return f
}

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
	reqSrv *reqsrv.RequestService[R, W],
	cmdEx *cmdex.CrudCommandService[C],
	rLim *rlim.RateLimitService[string],
	ttl *ttlserv.TimeToLiveService[time.Time],
)

func NewRateLimiter[
	C crud.CrudCommand,
	R reqlimprod.RateLimitProducer,
	S ~string,
	W req.Writer[S],
](path string) RateLimFunc[C, R, S, W] {
	f := func(
		reqSrv *reqsrv.RequestService[R, W],
		cmdEx *cmdex.CrudCommandService[C],
		rLim *rlim.RateLimitService[string],
		ttl *ttlserv.TimeToLiveService[time.Time],
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

func DefaultCfg[
	B any,
	C crud.CrudCommand,
	R CrudReq[B, C],
	S ~string,
	W req.Writer[S],
]() {
	f := func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService[C],
		rLim rlim.RateLimitService[string],
		ttl ttlserv.TimeToLiveService[time.Time],
	) {
		handle := func(r R, w W) {
			user, date, body := r.From(), r.Date(), r.Body()
			cmd, err := r.ProduceCmd()
			if err != nil {
				_ = w.Write(S(err.Error()))
				return
			}

		}
	}
}

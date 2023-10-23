package crud_controller

import (
	cmdex "in_mem_storage/application/service/command_executor"
	rlim "in_mem_storage/application/service/rate_limiter"
	reqsrv "in_mem_storage/application/service/request_handler"
	ttlserv "in_mem_storage/application/service/time_to_live"
	req "in_mem_storage/domain/incoming_request/abstraction"
	"in_mem_storage/domain/transaction/command/crud"
)

type (
// CommandExec = cmdex.CommandService[string, string, time.Time, time.Duration]
// RateLim     = rlim.RateLimitService[string, time.Time, time.Duration]
// ReqHandler  = reqsrv.RequestService[*http.Request, http.ResponseWriter]
)

type CrudReq[Body any, Command crud.CrudCommand] interface {
	req.Request[Body]
	crud.CrudCommandProducer[Command]
}

type CrudController[
	Body any,
	Command crud.CrudCommand,
	Read CrudReq[Body, Command],
	S ~string,
	Write req.Writer[S],
] struct {
	reqSrv reqsrv.RequestService[Read, Write]
	cmdEx  cmdex.CrudCommandService[Command]
	rLim   rlim.RateLimitService
	ttl    ttlserv.TimeToLiveService
}

func New[B any, C crud.CrudCommand, R CrudReq[B, C], S ~string, W req.Writer[S]](
	reqSrv reqsrv.RequestService[R, W],
	cmdEx cmdex.CrudCommandService[C],
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
) CrudController[B, C, R, S, W] {
	return CrudController[B, C, R, S, W]{reqSrv, cmdEx, rLim, ttl}
}

func (c *CrudController[B, C, R, S, W]) RunConfigs(
	cfgs ...func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService[C],
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
	),
) {
	for _, cfg := range cfgs {
		go cfg(c.reqSrv, c.cmdEx, c.rLim, c.ttl)
	}
}

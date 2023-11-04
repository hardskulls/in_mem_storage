package crud_controller

import (
	cmdex "in_mem_storage/internal/application/service/crud_cmd_executor"
	"in_mem_storage/internal/application/service/logger"
	rlim "in_mem_storage/internal/application/service/rate_limiter"
	reqsrv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	req "in_mem_storage/internal/domain/incoming_request/abstraction"
	crud "in_mem_storage/internal/domain/transaction/command/abstraction"
)

type CrudReq[Body any] interface {
	req.Request[Body]
	crud.CrudCommandProducer
}

type CrudController[Read, Write any] struct {
	reqSrv reqsrv.RequestService[Read, Write]
	cmdEx  cmdex.CrudCommandService
	rLim   rlim.RateLimitService
	ttl    ttlserv.TimeToLiveService
	logger logger.Logger
}

func New[Read, Write any](
	reqSrv reqsrv.RequestService[Read, Write],
	cmdEx cmdex.CrudCommandService,
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
	logger logger.Logger,
) CrudController[Read, Write] {
	return CrudController[Read, Write]{
		reqSrv,
		cmdEx,
		rLim,
		ttl,
		logger,
	}
}

func (c *CrudController[R, W]) RunConfig(
	cfg func(
		reqSrv reqsrv.RequestService[R, W],
		cmdEx cmdex.CrudCommandService,
		rLim rlim.RateLimitService,
		ttl ttlserv.TimeToLiveService,
		logger logger.Logger,
	),
) {
	cfg(c.reqSrv, c.cmdEx, c.rLim, c.ttl, c.logger)
}

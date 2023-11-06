package http_controller

import (
	cmdex "in_mem_storage/internal/application/service/crud_cmd_executor"
	"in_mem_storage/internal/application/service/logger"
	rlim "in_mem_storage/internal/application/service/rate_limiter"
	reqsrv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	reqabstr "in_mem_storage/internal/domain/incoming_request/abstraction"
	crud "in_mem_storage/internal/domain/transaction/command/abstraction"
)

const (
	BgGoroutineRunner ConfigRunner = iota
	MainGoroutineRunner
)

type CrudReq[Body any] interface {
	reqabstr.Request[Body]
	crud.CrudCommandProducer
}

type HttpController[Read, Write any] struct {
	reqSrv *reqsrv.RequestService[Read, Write]
	cmdEx  *cmdex.CrudCommandService
	rLim   *rlim.RateLimitService
	ttl    *ttlserv.TimeToLiveService
	logger *logger.Logger
}

type ConfigRunner int

func New[Read, Write any](
	reqSrv *reqsrv.RequestService[Read, Write],
	cmdEx *cmdex.CrudCommandService,
	rLim *rlim.RateLimitService,
	ttl *ttlserv.TimeToLiveService,
	logger *logger.Logger,
) HttpController[Read, Write] {
	return HttpController[Read, Write]{
		reqSrv: reqSrv,
		cmdEx:  cmdEx,
		rLim:   rLim,
		ttl:    ttl,
		logger: logger,
	}
}

func (c *HttpController[R, W]) RunConfigs(runner ConfigRunner, cfgs ...ReturningFunc[R, W]) {
	for _, cfg := range cfgs {
		switch runner {
		case BgGoroutineRunner:
			go cfg(c.reqSrv, c.cmdEx, c.rLim, c.ttl, c.logger)
		default:
			cfg(c.reqSrv, c.cmdEx, c.rLim, c.ttl, c.logger)
		}
	}
}

func (c *HttpController[R, W]) RunOn(port int) error {
	return c.reqSrv.RunServerOn(port)
}

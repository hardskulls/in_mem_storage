package http_controller

import (
	cmdex "in_mem_storage/internal/application/service/crud_cmd_executor"
	"in_mem_storage/internal/application/service/logger"
	rlim "in_mem_storage/internal/application/service/rate_limiter"
	reqsrv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	req "in_mem_storage/internal/domain/incoming_request/abstraction"
	crud "in_mem_storage/internal/domain/transaction/command/abstraction"
)

var (
	MainGoroutineRunner ConfigRunner = func(f func()) { f() }
	BgGoroutineRunner   ConfigRunner = func(f func()) { go f() }
)

type CrudReq[Body any] interface {
	req.Request[Body]
	crud.CrudCommandProducer
}

type HttpController[Read, Write any] struct {
	reqSrv  reqsrv.RequestService[Read, Write]
	cmdEx   cmdex.CrudCommandService
	rLim    rlim.RateLimitService
	ttl     ttlserv.TimeToLiveService
	logger  logger.Logger
	configs []ReturningFunc[Read, Write]
}

type ConfigRunner func(func())

func New[Read, Write any](
	reqSrv reqsrv.RequestService[Read, Write],
	cmdEx cmdex.CrudCommandService,
	rLim rlim.RateLimitService,
	ttl ttlserv.TimeToLiveService,
	logger logger.Logger,
) HttpController[Read, Write] {
	return HttpController[Read, Write]{
		reqSrv,
		cmdEx,
		rLim,
		ttl,
		logger,
		nil,
	}
}

func (c *HttpController[R, W]) RunConfig(
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

func (c *HttpController[R, W]) AppendConfigs(cfgs ...ReturningFunc[R, W]) {
	c.configs = append(c.configs, cfgs...)
}

func (c *HttpController[R, W]) RunConfigs(f ConfigRunner) {
	for _, cfg := range c.configs {
		config := func() { cfg(c.reqSrv, c.cmdEx, c.rLim, c.ttl, c.logger) }
		f(config)
	}
}

func (c *HttpController[R, W]) Run(port int) error {
	return c.reqSrv.Run(port)
}

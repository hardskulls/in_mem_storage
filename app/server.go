package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"in_mem_storage/internal/app/controller"
	"in_mem_storage/internal/app/server"
	"in_mem_storage/internal/implementation"
	"in_mem_storage/internal/service"
	"net/http"
	"time"
)

func initServer(port server.Port) *server.Server {
	srv := server.NewServer(port)

	srv = addMiddleware(srv)
	srv = addHandlers(srv)

	return srv
}

func addMiddleware(srv *server.Server) *server.Server {
	l := initLogger()

	rLimRepo := &implementation.RateLimRepo{}
	recRepo := &implementation.RecordRepo[string]{}
	ttlRepo := &implementation.TTLRepo{}

	cmdSrv := service.NewCommand(recRepo, ttlRepo)

	rLimCfg := controller.RateLimitConfig{Log: l, RLim: rLimRepo}
	cmdConfig := controller.CommandConfig{Log: l, Cmd: cmdSrv, RLim: rLimRepo}

	srv.Add(func(c *fiber.Ctx) error {
		c.Locals("RateLimitConfig", rLimCfg)
		c.Locals("CommandConfig", cmdConfig)

		return c.Next()
	})

	ttlConfig := controller.TimeToLiveConfig{Log: l, Cmd: cmdSrv, TTL: ttlRepo}
	go func() {
		ctx := context.Background()
		timout := time.Second
		user := "admin"

		_ = controller.TimeToLive(ctx, timout, user, ttlConfig)
	}()

	return srv
}

func addHandlers(srv *server.Server) *server.Server {

	srv.Handle(http.MethodPost, "/api/v1/ratelim", RateLimRoute)

	srv.Handle(http.MethodPost, "/api/v1/cmd", CmdRoute)

	return srv
}

func internalError(c *fiber.Ctx, err error) error {
	_ = c.
		Status(http.StatusInternalServerError).
		SendString("Internal Server Error")
	return err
}

func badRequest(c *fiber.Ctx, err error, msg string) error {
	_ = c.
		Status(http.StatusBadRequest).
		SendString(msg)
	return err
}

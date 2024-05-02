package controller

import (
	"context"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/data"
	"in_mem_storage/internal/domain/ratelim"
	"in_mem_storage/internal/domain/record"
	"in_mem_storage/internal/repository"
	"in_mem_storage/internal/service"
	"log/slog"
	"time"
)

type CommandConfig struct {
	RLim repository.RateLimit
	Cmd  service.Command
	Log  *slog.Logger
}

func Command(
	ctx context.Context,
	log *slog.Logger,
	user record.Author,
	cmd command.Command,
	rLim repository.RateLimit,
	cmdSrv service.Command,

) (data.JSON, error) {
	now := time.Now()

	rateLimit, err := rLim.GetFor(ctx, user)
	if err != nil {
		defaultZeroLim := ratelim.New(now, 0)
		rateLimit = defaultZeroLim
	}

	timeout := rateLimit.Timeout(now)

	time.Sleep(timeout)

	return cmdSrv.Execute(ctx, cmd)
}

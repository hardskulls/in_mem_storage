package controller

import (
	"context"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/data"
	"in_mem_storage/internal/domain/ratelim"
	"in_mem_storage/internal/domain/ttl"
	"in_mem_storage/internal/repository"
	"in_mem_storage/internal/service"
	"log/slog"
	"time"
)

type CommandConfig struct {
	RLim repository.RateLimit
	Cmd  service.Command
	TTL  repository.ExpiryCandidate
	Log  *slog.Logger
}

func Command(
	ctx context.Context,
	cmd command.Command,
	expires ttl.ExpirationTime,
	ec ttl.ExpiryCandidate,
	cfg CommandConfig,
) (data.JSON, error) {
	now := time.Now()

	rateLimit, err := cfg.RLim.GetFor(ctx, cmd.Author())
	if err != nil {
		defaultZeroLim := ratelim.New(now, 0)
		rateLimit = defaultZeroLim
	}

	timeout := rateLimit.Timeout(now)

	time.Sleep(timeout)

	if expires.After(time.Now()) && cmd.Type() == command.SetType {
		err := cfg.TTL.Set(ctx, expires, ec)
		if err != nil {
			return data.JSON{}, err
		}
	}

	return cfg.Cmd.Execute(ctx, cmd)
}

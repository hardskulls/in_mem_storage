package controller

import (
	"context"
	"in_mem_storage/internal/domain/ratelim"
	"in_mem_storage/internal/domain/record"
	"in_mem_storage/internal/repository"
	"log/slog"
)

type RateLimitConfig struct {
	RLim repository.RateLimit
	Log  *slog.Logger
}

func RateLimit(
	ctx context.Context,
	lim ratelim.RateLimit,
	user record.Author,
	cfg RateLimitConfig,
) error {
	err := cfg.RLim.SetFor(ctx, user, lim)
	if err != nil {
		return err
	}

	return nil
}

package controller

import (
	"context"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/record"
	"in_mem_storage/internal/repository"
	"in_mem_storage/internal/service"
	"log/slog"
	"time"
)

type TimeToLiveConfig struct {
	Cmd service.Command
	TTL repository.ExpiryCandidate
	Log *slog.Logger
}

func TimeToLive(
	ctx context.Context,
	log *slog.Logger,
	sleep time.Duration,
	user record.Author,
	with service.Command,
	ttl repository.ExpiryCandidate,
) error {
	for {
		time.Sleep(sleep)

		now := time.Now().Round(time.Second)
		expiredRec, err := ttl.Get(ctx, now)
		if err != nil {
			continue
		}

		cmd, err := command.NewCommand("delete", user, expiredRec, command.Empty{})
		if err != nil {
			continue
		}
		res, _ := with.Execute(ctx, cmd)

		log.Log(
			ctx,
			slog.LevelInfo,
			"A record expired.",
			slog.Any("key", expiredRec),
			slog.Any("result", res),
		)
	}
}

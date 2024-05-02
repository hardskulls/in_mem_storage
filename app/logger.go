package app

import (
	"log/slog"
	"os"
)

func initLogger() *slog.Logger {
	ops := &slog.HandlerOptions{Level: slog.LevelDebug}
	h := slog.NewTextHandler(os.Stdout, ops)

	return slog.New(h)
}

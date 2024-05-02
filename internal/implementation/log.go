package implementation

import (
	"context"
	"in_mem_storage/internal/domain/log"
	"log/slog"
)

type Log struct {
	l *slog.Logger
}

func NewLog(l *slog.Logger) Log {
	return Log{l: l}
}

func (l *Log) Save(ctx context.Context, entry log.Log[string]) error {
	l.l.Log(ctx, entry.LogLevel(), entry.Data(), entry.Event(), entry.Description())
	return nil
}

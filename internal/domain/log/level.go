package log

import "log/slog"

// Level is a value object representing a log level.
type Level = slog.Level

const (
	Error = slog.LevelError
	Info  = slog.LevelInfo
	Warn  = slog.LevelWarn
	Debug = slog.LevelDebug
)

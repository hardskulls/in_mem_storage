package log_record

import (
	fl "in_mem_storage/domain/log/value_object/file_line"
	"time"
)

type DefaultLogRecord = LogRecord[error]

type LogRecord[D any] struct {
	LogLevel        string
	Event           string
	Description     string
	Data            D
	EventOccurredAt time.Time
	FileLine        fl.FileLine
}

func New(lvl, event, descr string, err error) DefaultLogRecord {
	return DefaultLogRecord{
		LogLevel:        lvl,
		Event:           event,
		Description:     descr,
		Data:            err,
		EventOccurredAt: time.Now(),
		FileLine:        fl.New(2),
	}
}

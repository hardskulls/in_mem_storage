package log_record

import (
	"in_mem_storage/internal/domain/log/value_object"
	fl "in_mem_storage/internal/domain/log/value_object/file_line"
	stfrup "in_mem_storage/internal/domain/log/value_object/stack_frames_up"
	"time"
)

type DefaultLogRecord = LogRecord[string]

type LogRecord[D any] struct {
	logLevel    value_object.LogLvl
	event       value_object.Event
	description string
	data        D
	created     time.Time
	fileLine    fl.FileLine
}

func New(
	lvl value_object.LogLvl,
	event value_object.Event,
	descr, data string,
	stFramesUp stfrup.StackFramesUp,
) DefaultLogRecord {
	return DefaultLogRecord{
		logLevel:    lvl,
		event:       event,
		description: descr,
		data:        data,
		created:     time.Now(),
		fileLine:    fl.New(int(stFramesUp + stfrup.InOuterFn)),
	}
}

func (l LogRecord[D]) LogLevel() value_object.LogLvl {
	return l.logLevel
}

func (l LogRecord[D]) Event() value_object.Event {
	return l.event
}

func (l LogRecord[D]) Description() string {
	return l.description
}

func (l LogRecord[D]) Data() D {
	return l.data
}

func (l LogRecord[D]) Created() time.Time {
	return l.created
}

func (l LogRecord[D]) FileLine() fl.FileLine {
	return l.fileLine
}

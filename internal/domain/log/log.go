package log

import (
	"time"
)

type DefaultLogRecord = Log[string]

type CreatedAt = time.Time
type Description = string

// Log is a value object representing a log.
type Log[D any] struct {
	level       Level
	event       Event
	description Description
	data        D
	created     CreatedAt
	location    CodeLoc
}

// NewLog returns a Log instance.
func NewLog[D any](
	lvl Level,
	ev Event,
	d Description,
	data D,
	s Where,
) Log[D] {
	return Log[D]{
		level:       lvl,
		event:       ev,
		description: d,
		data:        data,
		created:     time.Now(),
		location:    NewCodeLoc(s + InOuterFn),
	}
}

func (l Log[D]) LogLevel() Level {
	return l.level
}

func (l Log[D]) Event() Event {
	return l.event
}

func (l Log[D]) Description() Description {
	return l.description
}

func (l Log[D]) Data() D {
	return l.data
}

func (l Log[D]) Created() CreatedAt {
	return l.created
}

func (l Log[D]) FileLine() CodeLoc {
	return l.location
}

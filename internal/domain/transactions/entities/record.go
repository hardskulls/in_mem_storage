package entities

import (
	"time"
)

type Record struct {
	Data             any
	CreatedAt        time.Time
	ExpiresInSeconds int
}

func NewRecord(data any, expiresIn int) Record {
	return Record{
		data,
		time.Now(),
		expiresIn,
	}
}

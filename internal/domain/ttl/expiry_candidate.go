package ttl

import (
	"in_mem_storage/internal/domain/record"
	"time"
)

type ExpirationTime = time.Time

type ExpiryCandidate struct {
	record record.ID
}

func New(r record.ID) ExpiryCandidate {
	return ExpiryCandidate{
		record: r,
	}
}

package implementation

import (
	"context"
	"errors"
	"in_mem_storage/internal/domain/record"
	"in_mem_storage/internal/domain/ttl"
	"sync"
)

type TTLRepo struct {
	inner sync.Map
}

func (t *TTLRepo) Set(
	ctx context.Context,
	created record.CreatedAt,
	ec ttl.ExpiryCandidate,
) error {
	t.inner.Store(created, ec)
	return nil
}

func (t *TTLRepo) Get(
	ctx context.Context,
	created record.CreatedAt,
) (ttl.ExpiryCandidate, error) {
	val, ok := t.inner.Load(created)
	if !ok {
		return ttl.ExpiryCandidate{}, errors.New("key not found")
	}
	return val.(ttl.ExpiryCandidate), nil
}

func (t *TTLRepo) Delete(
	ctx context.Context,
	created record.CreatedAt,
) error {
	t.inner.Delete(created)
	return nil
}

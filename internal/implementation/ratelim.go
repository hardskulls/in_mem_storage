package implementation

import (
	"context"
	"errors"
	"in_mem_storage/internal/domain/ratelim"
	"in_mem_storage/internal/domain/record"
	"sync"
)

type RateLimRepo struct {
	inner sync.Map
}

func (rl *RateLimRepo) GetFor(
	ctx context.Context,
	id record.ID,
) (ratelim.RateLimit, error) {
	limit, ok := rl.inner.Load(id)
	if !ok {
		return ratelim.RateLimit{}, errors.New("key not found")
	}

	return limit.(ratelim.RateLimit), nil
}
func (rl *RateLimRepo) SetFor(
	ctx context.Context,
	id record.ID,
	lim ratelim.RateLimit,
) error {
	rl.inner.Store(id, lim)
	return nil
}
func (rl *RateLimRepo) Remove(
	ctx context.Context,
	id record.ID,
) error {
	rl.inner.Delete(id)
	return nil
}

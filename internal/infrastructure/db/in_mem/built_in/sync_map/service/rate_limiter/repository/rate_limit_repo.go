package repository

import (
	lim "in_mem_storage/internal/domain/rate_limiter/value_object"
	"sync"
)

type RateLimitRepo[K comparable] struct {
	inner sync.Map
}

func New[K comparable]() RateLimitRepo[K] {
	return RateLimitRepo[K]{sync.Map{}}
}

func (r *RateLimitRepo[K]) Get(key K) (lim.RateLimit, error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return lim.RateLimit{}, GetError()
	}
	return val.(lim.RateLimit), nil
}

func (r *RateLimitRepo[K]) Set(key K, value lim.RateLimit) error {
	r.inner.Store(key, value)
	return nil
}

func (r *RateLimitRepo[K]) Delete(key K) error {
	r.inner.Delete(key)
	return nil
}

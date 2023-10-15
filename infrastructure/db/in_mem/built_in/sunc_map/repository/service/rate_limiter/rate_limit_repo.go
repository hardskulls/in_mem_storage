package rate_limiter

import (
	errs "in_mem_storage/domain/error"
	lim "in_mem_storage/domain/rate_limiter/value_object"
	"sync"
)

type RateLimitRepo[K comparable, U, L any] struct {
	inner sync.Map
}

func New[K comparable, U, L any]() RateLimitRepo[K, U, L] {
	return RateLimitRepo[K, U, L]{sync.Map{}}
}

func (r *RateLimitRepo[K, U, L]) Get(key K) (lim.RateLimit[U, L], errs.Error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return lim.RateLimit[U, L]{}, GetError()
	}
	return val.(lim.RateLimit[U, L]), errs.Error{}
}

func (r *RateLimitRepo[K, U, L]) Set(key K, value lim.RateLimit[U, L]) errs.Error {
	r.inner.Store(key, value)
	return errs.Error{}
}

func (r *RateLimitRepo[K, U, L]) Delete(key K) errs.Error {
	r.inner.Delete(key)
	return errs.Error{}
}

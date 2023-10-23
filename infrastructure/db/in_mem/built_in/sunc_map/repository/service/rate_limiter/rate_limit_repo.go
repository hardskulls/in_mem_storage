package rate_limiter

import (
	errs "in_mem_storage/domain/error"
	lim "in_mem_storage/domain/rate_limiter/value_object"
	"sync"
)

type RateLimitRepo[K comparable] struct {
	inner sync.Map
}

func New[K comparable]() RateLimitRepo[K] {
	return RateLimitRepo[K]{sync.Map{}}
}

func (r *RateLimitRepo[K]) Get(key K) (lim.RateLimit, errs.Error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return lim.RateLimit{}, GetError()
	}
	return val.(lim.RateLimit), errs.Error{}
}

func (r *RateLimitRepo[K]) Set(key K, value lim.RateLimit) errs.Error {
	r.inner.Store(key, value)
	return errs.Error{}
}

func (r *RateLimitRepo[K]) Delete(key K) errs.Error {
	r.inner.Delete(key)
	return errs.Error{}
}

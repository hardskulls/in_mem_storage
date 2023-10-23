package time_to_live

import (
	errs "in_mem_storage/domain/error"
	ec "in_mem_storage/domain/time_to_live/value_object"
	"sync"
)

type RateLimitRepo[K comparable] struct {
	inner sync.Map
}

func New[K comparable]() RateLimitRepo[K] {
	return RateLimitRepo[K]{sync.Map{}}
}

func (r *RateLimitRepo[K]) Get(key K) (ec.ExpiryCandidate, errs.Error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return ec.ExpiryCandidate{}, GetError()
	}
	return val.(ec.ExpiryCandidate), errs.Error{}
}

func (r *RateLimitRepo[K]) Set(key K, value ec.ExpiryCandidate) errs.Error {
	r.inner.Store(key, value)
	return errs.Error{}
}

func (r *RateLimitRepo[K]) Delete(key K) errs.Error {
	r.inner.Delete(key)
	return errs.Error{}
}

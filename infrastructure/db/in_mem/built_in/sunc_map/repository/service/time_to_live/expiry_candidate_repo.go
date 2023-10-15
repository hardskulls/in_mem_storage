package time_to_live

import (
	errs "in_mem_storage/domain/error"
	ec "in_mem_storage/domain/time_to_live/value_object"
	"sync"
)

type RateLimitRepo[K, C comparable] struct {
	inner sync.Map
}

func New[K, C comparable]() RateLimitRepo[K, C] {
	return RateLimitRepo[K, C]{sync.Map{}}
}

func (r *RateLimitRepo[K, C]) Get(key K) (ec.ExpiryCandidate[C], errs.Error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return ec.ExpiryCandidate[C]{}, GetError()
	}
	return val.(ec.ExpiryCandidate[C]), errs.Error{}
}

func (r *RateLimitRepo[K, C]) Set(key K, value ec.ExpiryCandidate[C]) errs.Error {
	r.inner.Store(key, value)
	return errs.Error{}
}

func (r *RateLimitRepo[K, C]) Delete(key K) errs.Error {
	r.inner.Delete(key)
	return errs.Error{}
}

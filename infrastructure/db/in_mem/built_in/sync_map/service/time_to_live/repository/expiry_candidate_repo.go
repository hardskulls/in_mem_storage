package repository

import (
	ec "in_mem_storage/domain/time_to_live/value_object"
	"sync"
)

type ExpiryRecRepo[K comparable] struct {
	inner sync.Map
}

func New[K comparable]() ExpiryRecRepo[K] {
	return ExpiryRecRepo[K]{sync.Map{}}
}

func (r *ExpiryRecRepo[K]) Get(key K) (ec.ExpiryRec, error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return ec.ExpiryRec{}, GetError()
	}
	return val.(ec.ExpiryRec), nil
}

func (r *ExpiryRecRepo[K]) Set(key K, value ec.ExpiryRec) error {
	r.inner.Store(key, value)
	return nil
}

func (r *ExpiryRecRepo[K]) Delete(key K) error {
	r.inner.Delete(key)
	return nil
}

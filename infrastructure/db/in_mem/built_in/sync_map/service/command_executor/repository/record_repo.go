package repository

import (
	errs "in_mem_storage/domain/error/value_object"
	rec "in_mem_storage/domain/transaction/record/value_object"
	"sync"
)

type RecordRepo[K comparable] struct {
	inner sync.Map
}

func New[K comparable]() RecordRepo[K] {
	return RecordRepo[K]{sync.Map{}}
}

func (r *RecordRepo[K]) Get(key K) (rec.Record, error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return rec.Record{}, GetError()
	}
	return val.(rec.Record), nil
}

func (r *RecordRepo[K]) Set(key K, value rec.Record) error {
	r.inner.Store(key, value)
	return nil
}

func (r *RecordRepo[K]) Delete(key K) error {
	r.inner.Delete(key)
	return nil
}

func (r *RecordRepo[K]) Update(key K, value rec.Record) error {
	_, err := r.Get(key)
	if err != (errs.Error{}) {
		return err
	}
	return r.Set(key, value)
}

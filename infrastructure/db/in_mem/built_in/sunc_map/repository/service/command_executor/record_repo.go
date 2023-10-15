package command_executor

import (
	errs "in_mem_storage/domain/error"
	rec "in_mem_storage/domain/transaction/record/value_object"
	"sync"
)

type RecordRepo[K comparable, D, T, E any] struct {
	inner sync.Map
}

func New[K comparable, D, T, E any]() RecordRepo[K, D, T, E] {
	return RecordRepo[K, D, T, E]{sync.Map{}}
}

func (r *RecordRepo[K, D, T, E]) Get(key K) (rec.Record[D, T, E], errs.Error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return rec.Record[D, T, E]{}, GetError()
	}
	return val.(rec.Record[D, T, E]), errs.Error{}
}

func (r *RecordRepo[K, D, T, E]) Set(key K, value rec.Record[D, T, E]) errs.Error {
	r.inner.Store(key, value)
	return errs.Error{}
}

func (r *RecordRepo[K, D, T, E]) Delete(key K) errs.Error {
	r.inner.Delete(key)
	return errs.Error{}
}

func (r *RecordRepo[K, D, T, E]) Update(key K, value rec.Record[D, T, E]) errs.Error {
	_, err := r.Get(key)
	if err != (errs.Error{}) {
		return err
	}
	return r.Set(key, value)
}

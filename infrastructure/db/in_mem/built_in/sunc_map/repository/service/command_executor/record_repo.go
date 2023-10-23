package command_executor

import (
	errs "in_mem_storage/domain/error"
	rec "in_mem_storage/domain/transaction/record/value_object"
	"sync"
)

type RecordRepo[K comparable] struct {
	inner sync.Map
}

func New[K comparable]() RecordRepo[K] {
	return RecordRepo[K]{sync.Map{}}
}

func (r *RecordRepo[K]) Get(key K) (rec.Record, errs.Error) {
	val, ok := r.inner.Load(key)
	if !ok {
		return rec.Record{}, GetError()
	}
	return val.(rec.Record), errs.Error{}
}

func (r *RecordRepo[K]) Set(key K, value rec.Record) errs.Error {
	r.inner.Store(key, value)
	return errs.Error{}
}

func (r *RecordRepo[K]) Delete(key K) errs.Error {
	r.inner.Delete(key)
	return errs.Error{}
}

func (r *RecordRepo[K]) Update(key K, value rec.Record) errs.Error {
	_, err := r.Get(key)
	if err != (errs.Error{}) {
		return err
	}
	return r.Set(key, value)
}

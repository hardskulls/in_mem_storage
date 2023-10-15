package standard_sync_map

import (
	errs "in_mem_storage/domain/errors"
	rec "in_mem_storage/domain/transactions/record/value_objects"
	"time"
)

type Rec = rec.Record[string, time.Time, time.Duration]

func (s *SyncMapStorage[K]) Get(key K) (Rec, errs.Error) {
	val, ok := s.inner.Load(key)
	if !ok {
		return Rec{}, GetError()
	}
	return val.(Rec), errs.Error{}
}

func (s *SyncMapStorage[K]) Set(key K, value Rec) errs.Error {
	s.inner.Store(key, value)
	return errs.Error{}
}

func (s *SyncMapStorage[K]) Delete(key K) errs.Error {
	_, existed := s.inner.LoadAndDelete(key)
	if !existed {
		return DeleteError()
	}
	return errs.Error{}
}

func (s *SyncMapStorage[K]) Update(key K, value Rec) errs.Error {
	_, err := s.Get(key)

	if (errs.Error{}) != err {
		return err
	}

	return s.Set(key, value)
}

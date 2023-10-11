package standard_sync_map

import (
	rec "in_mem_storage/internal/domain/transactions/entities"
	errs "in_mem_storage/internal/domain/errors"
)

func (s *SyncMapStorage) GetValue(key any) (rec.Record, errs.Error) {
	val, ok := s.Load(key)
	if !ok {
		return rec.Record{}, GetError()
	}
	return val.(rec.Record), errs.Error{}
}

func (s *SyncMapStorage) SetValue(key any, value rec.Record) errs.Error {
	s.Store(key, value)
	return errs.Error{}
}

func (s *SyncMapStorage) DeleteValue(key any) errs.Error {
	_, existed := s.LoadAndDelete(key)
	if !existed {
		return DeleteError()
	}
	return errs.Error{}
}

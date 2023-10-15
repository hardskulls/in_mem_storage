package standard_sync_map

import (
	errs "in_mem_storage/domain/errors"
	er "in_mem_storage/domain/time_to_live/value_objects"
)

type ExpRec = er.ExpiryRecord[string]

func (s *SyncMapStorage[K]) AddExpiryRecord(key K, value ExpRec) errs.Error {
	s.inner.Store(key, value)
	return errs.Error{}
}

func (s *SyncMapStorage[K]) GetExpiryRecord(key K, value ExpRec) errs.Error {
	s.inner.Store(key, value)
	return errs.Error{}
}

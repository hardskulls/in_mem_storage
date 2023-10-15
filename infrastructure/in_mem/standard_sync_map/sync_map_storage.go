package standard_sync_map

import (
	"sync"
)

type SyncMapStorage[K comparable] struct {
	inner sync.Map
}

// func(*SyncMapStorage) New() SyncMapStorage {
// 	return SyncMapStorage{sync.Map{}}
// }

// func (s *SyncMapStorage) Set(key t.Key, value t.Value) error {
// 	s.Store(key, value)
// 	return nil
// }

// func (s *SyncMapStorage) Get(key t.Key) (t.Value, error) {
// 	any, ok := s.Load(key)
// 	if !ok {
// 		return "", GetError()
// 	}
// 	return any.(t.Value), nil
// }

// func (s *SyncMapStorage) Delete(key t.Key) error {
// 	_, existed := s.LoadAndDelete(key)
// 	if !existed {
// 		return DeleteError()
// 	}
// 	return nil
// }

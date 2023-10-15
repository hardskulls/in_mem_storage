package repository

import (
	er "in_mem_storage/domain/time_to_live/value_objects"
)

type ExpiryRecordRepo[K, R comparable] interface {
	Add(key K, value er.ExpiryRecord[R]) error
	Get(key K) (er.ExpiryRecord[R], error)
	Remove(key K) error
}

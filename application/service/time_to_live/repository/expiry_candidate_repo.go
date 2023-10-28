package repository

import (
	er "in_mem_storage/domain/time_to_live/value_object"
	"time"
)

type DefaultExpiryRecRepo = ExpiryRecRepo[time.Time]

type ExpiryRecRepo[K comparable] interface {
	Set(key K, value er.ExpiryRec) error
	Get(key K) (er.ExpiryRec, error)
	Delete(key K) error
}

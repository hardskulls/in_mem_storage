package repositories

import (
	rec "in_mem_storage/internal/domain/transactions/entities"
)

type RecordRepo[K comparable] interface {
	GetValue(key K) (rec.Record, error)
	SetValue(key K, value rec.Record) error
	DeleteValue(key K) error
}

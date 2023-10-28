package repository

import (
	rec "in_mem_storage/domain/transaction/record/value_object"
)

type DefaultRecordRepo = RecordRepo[string]

type RecordRepo[K comparable] interface {
	Set(key K, value rec.Record) error
	Get(key K) (rec.Record, error)
	Update(key K, value rec.Record) error
	Delete(key K) error
}

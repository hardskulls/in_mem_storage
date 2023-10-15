package repository

import (
	rec "in_mem_storage/domain/transaction/record/value_object"
)

type RecordRepo[K comparable, D, T, E any] interface {
	Set(key K, value rec.Record[D, T, E]) error
	Get(key K) (rec.Record[D, T, E], error)
	Update(key K, value rec.Record[D, T, E]) error
	Delete(key K) error
}

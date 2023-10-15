package value_objects

import (
	rec "in_mem_storage/domain/transactions/record/value_objects"
)

type UpdateCommand[K comparable, D, T, E any] struct {
	key   K
	value rec.Record[D, T, E]
}

func NewUpdateCommand[K comparable, D, T, E any](key K, value rec.Record[D, T, E]) UpdateCommand[K, D, T, E] {
	return UpdateCommand[K, D, T, E]{key, value}
}

func (d UpdateCommand[K, _, _, _]) Key() K {
	return d.key
}

func (d UpdateCommand[_, D, T, E]) Value() rec.Record[D, T, E] {
	return d.value
}

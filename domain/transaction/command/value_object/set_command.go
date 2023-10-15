package value_objects

import (
	rec "in_mem_storage/domain/transactions/record/value_objects"
)

type SetCommand[K comparable, D, T, E any] struct {
	key   K
	value rec.Record[D, T, E]
}

func NewSetCommand[K comparable, D, T, E any](key K, value rec.Record[D, T, E]) SetCommand[K, D, T, E] {
	return SetCommand[K, D, T, E]{key, value}
}

func (d SetCommand[K, _, _, _]) Key() K {
	return d.key
}

func (d SetCommand[_, D, T, E]) Value() rec.Record[D, T, E] {
	return d.value
}

package value_objects

import (
	rec "in_mem_storage/internal/domain/transactions/entities"
	repo "in_mem_storage/internal/domain/transactions/repositories"
)

type SetCommand[K comparable, V any] struct {
	Key       K
	Value     V
	ExpiresIn int
}

func recFromSetCmd[K comparable, V any](cmd SetCommand[K, V]) rec.Record {
	return rec.NewRecord(cmd.Value, cmd.ExpiresIn)
}

func (c *SetCommand[K, V]) Execute(r interface{ repo.RecordRepo[K] }) error {
	record := recFromSetCmd[K, V](*c)
	err := r.SetValue(c.Key, record)
	return err
}

// -----------------------------------------------------------

// type Foo[T, A any] interface {
// 	Foo(t T) A
// }

// type Bar[T any] struct {
// }

// type R struct {
// }

// func (b *Bar[T]) Foo(t any) string {
// 	return "rec"
// }

// func bbbz() {
// 	bar := Bar[string]{}
// 	foo := bar.Foo("foo")
// 	print(foo)
// }

// -----------------------------------------------------------

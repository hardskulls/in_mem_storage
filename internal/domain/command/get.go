package command

import "in_mem_storage/internal/domain/record"

type Get[T comparable] struct {
	author record.Author
	key    T
}

func (c *Get[K]) Author() record.Author {
	return c.author
}

func (c *Get[T]) Key() T {
	return c.key
}

func (c *Get[T]) IsCommand() {}

func (c *Get[K]) Type() Type {
	return GetType
}

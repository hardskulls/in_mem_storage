package command

import "in_mem_storage/internal/domain/record"

type Set[K comparable, V any] struct {
	author record.Author
	key    K
	value  V
}

func (c *Set[K, V]) Author() record.Author {
	return c.author
}

func (c *Set[K, V]) Key() K {
	return c.key
}

func (c *Set[K, V]) Value() V {
	return c.value
}

func (c *Set[K, V]) IsCommand() {}

func (c *Set[K, V]) Type() Type {
	return SetType
}

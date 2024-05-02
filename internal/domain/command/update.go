package command

import "in_mem_storage/internal/domain/record"

type Update[K comparable, V any] struct {
	author record.Author
	key    K
	value  V
}

func (c *Update[K, V]) Author() record.Author {
	return c.author
}

func (c *Update[K, V]) Key() K {
	return c.key
}

func (c *Update[K, V]) Value() V {
	return c.value
}

func (c *Update[K, V]) IsCommand() {}

func (c *Update[K, V]) Type() Type {
	return UpdateType
}

package command

import "in_mem_storage/internal/domain/record"

type Delete[K comparable] struct {
	author record.Author
	key    K
}

func (c *Delete[K]) Author() record.Author {
	return c.author
}

func (c *Delete[K]) Key() K {
	return c.key
}

func (c *Delete[K]) IsCommand() {}

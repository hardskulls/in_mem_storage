package value_objects

import (
// "time"
)

type Record[D, T, E any] struct {
	data      D
	createdAt T
	expiresIn E
}

func NewRecord[D, T, E any](data D, createdAt T, expiresIn E) Record[D, T, E] {
	return Record[D, T, E]{
		data,
		createdAt,
		expiresIn,
	}
}

func (r Record[D, _, _]) Data() D {
	return r.data
}

func (r Record[_, T, _]) CreatedAt() T {
	return r.createdAt
}

func (r Record[_, _, E]) ExpiresIn() E {
	return r.expiresIn
}

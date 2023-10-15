package value_objects

type DeleteCommand[K comparable] struct {
	key K
}

func NewDeleteCommand[K comparable](key K) DeleteCommand[K] {
	return DeleteCommand[K]{key}
}

func (d DeleteCommand[K]) Key() K {
	return d.key
}

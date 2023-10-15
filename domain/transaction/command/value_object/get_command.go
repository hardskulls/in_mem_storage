package value_objects

type GetCommand[K comparable] struct {
	key K
}

func NewGetCommand[K comparable](key K) GetCommand[K] {
	return GetCommand[K]{key}
}

func (d GetCommand[K]) Key() K {
	return d.key
}

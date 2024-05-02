package convert

type TryInto[T any] interface {
	TryInto() (T, error)
}

package errors

type Result[T, E any] struct {
	Ok  T
	Err E
}

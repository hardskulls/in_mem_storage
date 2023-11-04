package value_object

type Result[T, E any] struct {
	Ok  T
	Err E
}

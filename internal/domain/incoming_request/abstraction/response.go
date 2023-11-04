package abstraction

type Writer[D any] interface {
	Write(D) error
}

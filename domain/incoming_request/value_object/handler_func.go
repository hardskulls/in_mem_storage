package value_objects

type ReqHandlerFunc[R, W any] struct {
	Path   string
	Handle func(req R, resp W)
}

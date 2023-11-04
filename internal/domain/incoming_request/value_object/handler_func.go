package value_objects

type ReqHandler[Read, Write any] struct {
	Path   string
	Handle func(req Read, resp Write)
}

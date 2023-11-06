package abstraction

type Request[Body any] interface {
	From() string
	Body() (Body, error)
}

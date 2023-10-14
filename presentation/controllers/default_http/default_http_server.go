package default_http

import (
	"io"
	"net/http"
	"in_mem_storage/internal/application/services"
)

type DefaultHttp struct {
}

func (s *DefaultHttp) Handle(req *http.Request, resp io.Writer) {
	http.HandleFunc()
}

type CommandController[Req, Resp any] struct {
	reqHandler services.RequestHandler[Req, Resp]
	recordRepo  any
}

func (c *CommandController[Req, Resp]) Do(f func(req *http.Request, resp io.Writer), handler Htt) {
	http.HandleFunc()
}

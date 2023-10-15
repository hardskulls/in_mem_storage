package default_http

import (
	"in_mem_storage/application/services"
	"io"
	"net/http"
)

type DefaultHttp struct {
}

func (s *DefaultHttp) Handle(req *http.Request, resp io.Writer) {
	http.HandleFunc()
}

type CommandController[Req, Resp any] struct {
	reqHandler services.RequestHandler[Req, Resp]
	recordRepo any
}

func (c *CommandController[Req, Resp]) Do(f func(req *http.Request, resp io.Writer), handler Htt) {
	http.HandleFunc()
}

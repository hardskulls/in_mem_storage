package request

import (
	"fmt"
	errs "in_mem_storage/internal/domain/error/value_object"
	"net/http"
)

func MissingParamError(missing string) error {
	return errs.New(fmt.Sprintf("[RequestBodyError] : missing required parameter: %v", missing), 1)
}

type Request struct {
	inner *http.Request
}

func New(r *http.Request) Request {
	return Request{inner: r}
}

func (r Request) Body() string {
	var buff []byte
	_, _ = r.inner.Body.Read(buff)
	return string(buff)
}

func (r Request) From() string {
	return r.inner.RemoteAddr
}

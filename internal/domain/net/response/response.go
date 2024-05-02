package response

import "net/http"

type Writer[D any] interface {
	Write(D) error
}

type Response struct {
	inner http.ResponseWriter
}

func New(w http.ResponseWriter) Response {
	return Response{inner: w}
}

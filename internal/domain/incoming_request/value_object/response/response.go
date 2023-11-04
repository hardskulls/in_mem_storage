package response

import "net/http"

type Response struct {
	inner http.ResponseWriter
}

func New(w http.ResponseWriter) Response {
	return Response{inner: w}
}

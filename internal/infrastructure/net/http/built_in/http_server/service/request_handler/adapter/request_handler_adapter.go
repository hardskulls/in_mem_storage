package adapter

import (
	req "in_mem_storage/internal/domain/incoming_request/value_object"
	"in_mem_storage/internal/domain/incoming_request/value_object/request"
	"in_mem_storage/internal/domain/incoming_request/value_object/response"
	"net/http"
	"strconv"
)

type ReqHandler = req.ReqHandler[request.Request, response.Response]

type StandardHTTPRequestAdapter struct{}

func (s StandardHTTPRequestAdapter) Handle(handler ReqHandler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		handler.Handle(request.New(r), response.New(w))
	}
	http.HandleFunc(handler.Path, f)
}

func (s StandardHTTPRequestAdapter) Run(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

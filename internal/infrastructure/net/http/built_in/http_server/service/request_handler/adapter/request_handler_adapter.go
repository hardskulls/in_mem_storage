package adapter

import (
	increq "in_mem_storage/internal/domain/incoming_request/value_object"
	"in_mem_storage/internal/domain/incoming_request/value_object/request"
	resp "in_mem_storage/internal/domain/incoming_request/value_object/response"
	"net/http"
	"strconv"
)

type StandardHTTPRequestAdapter struct{}

func (s *StandardHTTPRequestAdapter) Handle(handler increq.DefaultReqHandler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		handler.Handle(request.New(r), resp.New(w))
	}
	http.HandleFunc(handler.Path, f)
}

func (s *StandardHTTPRequestAdapter) Run(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

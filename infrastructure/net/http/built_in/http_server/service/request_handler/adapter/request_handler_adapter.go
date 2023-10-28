package adapter

import (
	req "in_mem_storage/domain/incoming_request/value_object"
	"net/http"
)

type ReqHandler = req.ReqHandler[*http.Request, http.ResponseWriter]

type StandardHTTPRequestAdapter struct {
}

func (s *StandardHTTPRequestAdapter) Handle(handler ReqHandler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		handler.Handle(r, w)
	}
	http.HandleFunc(handler.Path, f)
}

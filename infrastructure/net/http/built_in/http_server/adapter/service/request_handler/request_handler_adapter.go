package request_handler

import (
	req "in_mem_storage/domain/incoming_request/value_object"
	"net/http"
)

type ReqHandler = req.ReqHandlerFunc[*http.Request, http.ResponseWriter]

type StandardHTTPRequestAdapter struct {
}

func (s *StandardHTTPRequestAdapter) Handle(handlers ...ReqHandler) {
	for _, h := range handlers {
		f := func(w http.ResponseWriter, r *http.Request) {
			h.Handle(r, w)
		}
		go http.HandleFunc(h.Path, f)
	}
}

package port

import (
	req "in_mem_storage/internal/domain/incoming_request/value_object"
)

type ReqHandlerPort[R, W any] interface {
	Handle(handler req.ReqHandler[R, W])
	Run(port int) error
}

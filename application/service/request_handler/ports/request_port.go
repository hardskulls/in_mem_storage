package ports

import (
	req "in_mem_storage/domain/incoming_request/value_object"
)

type RequestPort[R, W any] interface {
	Handle(handlers ...req.ReqHandlerFunc[R, W])
}

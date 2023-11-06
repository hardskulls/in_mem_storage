package value_objects

import (
	"in_mem_storage/internal/domain/incoming_request/value_object/request"
	"in_mem_storage/internal/domain/incoming_request/value_object/response"
)

type DefaultReqHandler = ReqHandler[request.Request, response.Response]

type ReqHandler[Read, Write any] struct {
	Path   string
	Handle func(req Read, resp Write)
}

package services

import (
	"in_mem_storage/internal/application/service/server/port"
	req "in_mem_storage/internal/domain/incoming_request/value_object"
)

type RequestService[Read, Write any] struct {
	requestPort port.ReqHandlerPort[Read, Write]
}

func New[R, W any](with port.ReqHandlerPort[R, W]) RequestService[R, W] {
	return RequestService[R, W]{with}
}

func (rs *RequestService[R, W]) Handle(handler req.ReqHandler[R, W]) {
	rs.requestPort.Handle(handler)
}

func (rs *RequestService[R, W]) Run(port int) error {
	return rs.requestPort.Run(port)
}

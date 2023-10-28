package services

import (
	port "in_mem_storage/application/service/server/port"
	req "in_mem_storage/domain/incoming_request/value_object"
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

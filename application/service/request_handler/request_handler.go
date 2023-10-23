package services

import (
	port "in_mem_storage/application/service/request_handler/port"
	req "in_mem_storage/domain/incoming_request/value_object"
)

type RequestService[Read, Write any] struct {
	requestPort port.RequestPort[Read, Write]
}

func New[R, W any](with port.RequestPort[R, W]) RequestService[R, W] {
	return RequestService[R, W]{with}
}

func (rs *RequestService[R, W]) Handle(handlers ...req.ReqHandler[R, W]) {
	rs.requestPort.Handle(handlers...)
}

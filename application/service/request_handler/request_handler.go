package services

import (
	port "in_mem_storage/application/service/request_handler/ports"
	req "in_mem_storage/domain/incoming_request/value_object"
)

type RequestService[R, W any] struct {
	requestPort port.RequestPort[R, W]
}

func NewRequestService[R, W any](with port.RequestPort[R, W]) RequestService[R, W] {
	return RequestService[R, W]{with}
}

func (rs *RequestService[R, W]) Handle(handlers ...req.ReqHandlerFunc[R, W]) {
	rs.requestPort.Handle(handlers...)
}

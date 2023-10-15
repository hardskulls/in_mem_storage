package services

import (
	port "in_mem_storage/application/service/request_handler/ports"
	req "in_mem_storage/domain/incoming_request/value_object"
)

type RequestService[D any] struct {
	requestPort port.RequestPort[D]
}

func NewRequestService[D any](with port.RequestPort[D]) RequestService[D] {
	return RequestService[D]{with}
}

func (rs *RequestService[D]) Handle(r func(req *req.Request[D])) {}

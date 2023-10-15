package services

// errs "in_mem_storage/domain/errors"
// ts "in_mem_storage/domain/transactions/abstractions"
// rec "in_mem_storage/domain/transactions/entities"
// events "in_mem_storage/domain/transactions/events"
// cmd "in_mem_storage/domain/transactions/value_objects"
// "net/http"

import (
	rp "in_mem_storage/domain/incoming_requests/ports"
	req "in_mem_storage/domain/incoming_requests/value_objects"
)

// type RequestHandler[REQ, RESP any] interface {
// 	Handle(f func(req REQ, resp RESP))
// }

type RequestService[D any] struct {
	requestPort rp.RequestPort[D]
}

func NewRequestService[D any](with rp.RequestPort[D]) RequestService[D] {
	return RequestService[D]{with}
}

func (rs *RequestService[D]) Handle(r func(req *req.Request[D])) {}

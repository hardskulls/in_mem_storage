package services

import (
	// errs "in_mem_storage/internal/domain/errors"
	// ts "in_mem_storage/internal/domain/transactions/abstractions"
	// rec "in_mem_storage/internal/domain/transactions/entities"
	// events "in_mem_storage/internal/domain/transactions/events"
	// cmd "in_mem_storage/internal/domain/transactions/value_objects"
	// "net/http"
)

type RequestHandler[REQ, RESP any] interface {
	Handle(f func(req REQ, resp RESP))
}

package net

import (
	"in_mem_storage/internal/domain/net/request"
	"in_mem_storage/internal/domain/net/response"
)

type DefaultReqHandler = ReqHandler[request.BasicRequest, response.Response]

type ReqHandler[Read, Write any] struct {
	Path   string
	Handle func(req Read, resp Write)
}

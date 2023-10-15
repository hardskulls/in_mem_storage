package ports

import (
	req "in_mem_storage/domain/incoming_request/value_object"
)

type RequestPort[D any] interface {
	Handle(f func(r req.Request[D]))
}

package ports

import (
	req "in_mem_storage/domain/incoming_requests/value_objects"
)

type RequestPort[D any] interface {
	Handle(f func(r req.Request[D]))
}

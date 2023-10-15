package repositories

import (
	lim "in_mem_storage/domain/rate_limiter/value_objects"
)

type RateLimitRepo[U, L any, K comparable] interface {
	Get(by K) (lim.RateLimit[U, L], error)
}

package repositories

import (
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitRepo[U, L any, K comparable] interface {
	Get(by K) (lim.RateLimit[U, L], error)
	Set(key K, value lim.RateLimit[U, L]) error
	Delete(key K) error
}

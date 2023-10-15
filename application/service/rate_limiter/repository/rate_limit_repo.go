package repositories

import (
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitRepo[K comparable, U, L any] interface {
	Get(by K) (lim.RateLimit[U, L], error)
	Set(key K, value lim.RateLimit[U, L]) error
	Delete(key K) error
}

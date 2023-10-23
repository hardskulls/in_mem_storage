package repositories

import (
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitRepo[K comparable] interface {
	Get(key K) (lim.RateLimit, error)
	Set(key K, value lim.RateLimit) error
	Delete(key K) error
}

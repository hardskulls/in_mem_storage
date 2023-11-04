package repositories

import (
	lim "in_mem_storage/internal/domain/rate_limiter/value_object"
)

type DefaultRateLimitRepo = RateLimitRepo[string]

type RateLimitRepo[K comparable] interface {
	Get(key K) (lim.RateLimit, error)
	Set(key K, value lim.RateLimit) error
	Delete(key K) error
}

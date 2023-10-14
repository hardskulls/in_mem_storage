package repositories

import (
	lim "in_mem_storage/internal/domain/rate_limiter/value_objects"
)

type RateLimitRepo interface {
	Get(by any) (lim.RateLimit, error)
}

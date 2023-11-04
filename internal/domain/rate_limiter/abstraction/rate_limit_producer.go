package abstraction

import (
	lim "in_mem_storage/internal/domain/rate_limiter/value_object"
)

type RateLimitProducer interface {
	ProduceRateLim() (lim.RateLimit, error)
}

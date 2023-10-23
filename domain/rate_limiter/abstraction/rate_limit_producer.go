package abstraction

import lim "in_mem_storage/domain/rate_limiter/value_object"

type RateLimitProducer interface {
	Produce() (lim.RateLimit, error)
}

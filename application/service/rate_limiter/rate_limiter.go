package services

import (
	repo "in_mem_storage/application/service/rate_limiter/repository"
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitService[Key comparable] struct {
	rateLimitRepo repo.RateLimitRepo[Key]
}

func WithRateLimitRepo[D any, K comparable](repo repo.RateLimitRepo[K]) RateLimitService[K] {
	return RateLimitService[K]{rateLimitRepo: repo}
}

func (rl *RateLimitService[K]) Get(key K) (lim.RateLimit, error) {
	return rl.rateLimitRepo.Get(key)
}

func (rl *RateLimitService[K]) Set(key K, value lim.RateLimit) error {
	return rl.rateLimitRepo.Set(key, value)
}

func (rl *RateLimitService[K]) Delete(key K) error {
	return rl.rateLimitRepo.Delete(key)
}

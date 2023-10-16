package services

import (
	repo "in_mem_storage/application/service/rate_limiter/repository"
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitService[K comparable, U, L any] struct {
	rateLimitRepo repo.RateLimitRepo[K, U, L]
}

func WithRateLimitRepo[U, L any, K comparable](repo repo.RateLimitRepo[K, U, L]) RateLimitService[K, U, L] {
	return RateLimitService[K, U, L]{rateLimitRepo: repo}
}

func (rl *RateLimitService[K, U, L]) Get(key K) (lim.RateLimit[U, L], error) {
	return rl.rateLimitRepo.Get(key)
}

func (rl *RateLimitService[K, U, L]) Set(key K, value lim.RateLimit[U, L]) error {
	return rl.rateLimitRepo.Set(key, value)
}

func (rl *RateLimitService[K, U, L]) Delete(key K) error {
	return rl.rateLimitRepo.Delete(key)
}

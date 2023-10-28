package services

import (
	repo "in_mem_storage/application/service/rate_limiter/repository"
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitService struct {
	rateLimitRepo repo.DefaultRateLimitRepo
}

func New(repo repo.DefaultRateLimitRepo) RateLimitService {
	return RateLimitService{rateLimitRepo: repo}
}

func (rl *RateLimitService) Get(key string) (lim.RateLimit, error) {
	return rl.rateLimitRepo.Get(key)
}

func (rl *RateLimitService) Set(key string, value lim.RateLimit) error {
	return rl.rateLimitRepo.Set(key, value)
}

func (rl *RateLimitService) Delete(key string) error {
	return rl.rateLimitRepo.Delete(key)
}

package services

import (
	repo "in_mem_storage/application/service/rate_limiter/repository"
	lim "in_mem_storage/domain/rate_limiter/value_object"
)

type RateLimitService[U, L any, K comparable] struct {
	rateLimitRepo repo.RateLimitRepo[U, L, K]
}

func WithRateLimitRepo[U, L any, K comparable](repo repo.RateLimitRepo[U, L, K]) RateLimitService[U, L, K] {
	return RateLimitService[U, L, K]{rateLimitRepo: repo}
}

func (rl *RateLimitService[U, L, K]) Get(key K) (lim.RateLimit[U, L], error) {
	return rl.rateLimitRepo.Get(key)
}

func (rl *RateLimitService[U, L, K]) Set(key K, value lim.RateLimit[U, L]) error {
	return rl.rateLimitRepo.Set(key, value)
}

func (rl *RateLimitService[U, L, K]) Delete(key K) error {
	return rl.rateLimitRepo.Delete(key)
}

// type TokenBucket struct {
// 	*sync.Mutex
// 	startedAt     time.Time
// 	maxTokens     int
// 	currentTokens int
// 	refillAfter   time.Duration
// }

// func (tb *TokenBucket) lockDoUnlock(f func()) {
// 	tb.Lock()
// 	f()
// 	tb.Unlock()
// }

// func NewTokenBucket(tokens int, refillAfter time.Duration) TokenBucket {
// 	b := TokenBucket{
// 		startedAt:     time.Now(),
// 		maxTokens:     tokens,
// 		currentTokens: tokens,
// 		refillAfter:   refillAfter,
// 	}
// 	return b
// }

// func (tb *TokenBucket) AvailableTokens(consumer string, spend int) int {
// 	var currentTokens int
// 	tb.lockDoUnlock(func() {
// 		currentTime := time.Now()
// 		timeToRefill := tb.startedAt.Add(tb.refillAfter)
// 		if currentTime.Compare(timeToRefill) >= 0 {
// 			tb.currentTokens = tb.maxTokens
// 		}
// 		tb.currentTokens -= spend
// 		currentTokens = tb.currentTokens
// 	})
// 	return currentTokens
// }

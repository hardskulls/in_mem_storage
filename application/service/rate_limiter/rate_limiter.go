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

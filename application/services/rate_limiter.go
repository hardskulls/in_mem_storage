package services

import (
	// rl "in_mem_storage/internal/domain/rate_limiter/abstractions"
	// "sync"
	// "time"
	repo "in_mem_storage/internal/domain/rate_limiter/repositories"
	lim "in_mem_storage/internal/domain/rate_limiter/value_objects"
)

type RateLimiter struct {
	rateLimitRepo repo.RateLimitRepo
}

func WithRateLimitRepo(repo repo.RateLimitRepo) RateLimiter {
	return RateLimiter{rateLimitRepo: repo}
}

func (rl *RateLimiter) GetBy(by any) (lim.RateLimit, error) {
	return rl.rateLimitRepo.Get(by)
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

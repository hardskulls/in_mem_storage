package objects

import (
	"sync"
	"time"
)

type TokenBucket struct {
	*sync.Mutex
	startedAt     time.Time
	maxTokens     int
	currentTokens int
	refillAfter   time.Duration
}

func (tb *TokenBucket) safelyOperate(f func()) {
	tb.Lock()
	f()
	tb.Unlock()
}

func NewTokenBucket(tokens int, refillAfter time.Duration) TokenBucket {
	b := TokenBucket{
		startedAt:     time.Now(),
		maxTokens:     tokens,
		currentTokens: tokens,
		refillAfter:   refillAfter,
	}
	return b
}

func (r *TokenBucket) AvailableTokens(spend int) int {
	var currentTokens int
	r.safelyOperate(func() {
		currentTime := time.Now()
		timeToRefill := r.startedAt.Add(r.refillAfter)
		if currentTime.Compare(timeToRefill) >= 0 {
			r.currentTokens = r.maxTokens
		}
		r.currentTokens -= spend
		currentTokens = r.currentTokens
	})
	return currentTokens
}




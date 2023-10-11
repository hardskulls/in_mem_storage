package abstractions

import (
	// "time"
)

type RateLimiter interface {
	AvailableTokens(spend int) int
}

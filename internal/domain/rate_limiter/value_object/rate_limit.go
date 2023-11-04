package value_objects

import "time"

// import "time"

type RateLimit struct {
	For      string
	LastUsed time.Time
	Limit    time.Duration
}

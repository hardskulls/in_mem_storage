package ratelim

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

type Limit = time.Duration
type LastUsed = time.Time

// RateLimit as a value object.
type RateLimit struct {
	lastUsed LastUsed
	limit    Limit
}

func New(used LastUsed, lim Limit) RateLimit {
	return RateLimit{
		lastUsed: used,
		limit:    lim,
	}
}

func (r *RateLimit) Timeout(from time.Time) time.Duration {
	return r.lastUsed.Add(r.limit).Sub(from)
}

func (r *RateLimit) UpdateLastUsed() {
	r.lastUsed = time.Now()
}

func (r *RateLimit) SetLimit(d Limit) {
	r.limit = d
}

func FromRequest(r *http.Request) (RateLimit, error) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		return RateLimit{}, errors.New("no limit specified")
	}

	digit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return RateLimit{}, err
	}
	rateLim := time.Duration(digit) * time.Millisecond

	return RateLimit{
		lastUsed: time.Now(),
		limit:    rateLim,
	}, nil
}

package repository

import (
	"context"
	"in_mem_storage/internal/domain/ratelim"
	"in_mem_storage/internal/domain/record"
)

// RateLimit is a repository for RateLimit.
type RateLimit interface {
	GetFor(ctx context.Context, a record.Author) (ratelim.RateLimit, error)
	SetFor(ctx context.Context, a record.Author, lim ratelim.RateLimit) error
	Remove(ctx context.Context, a record.Author) error
}

package repository

import (
	"context"
	"in_mem_storage/internal/domain/ttl"
)

type ExpiryCandidate interface {
	Set(ctx context.Context, expires ttl.ExpirationTime, ec ttl.ExpiryCandidate) error
	Get(ctx context.Context, expires ttl.ExpirationTime) (ttl.ExpiryCandidate, error)
	Delete(ctx context.Context, expires ttl.ExpirationTime) error
}

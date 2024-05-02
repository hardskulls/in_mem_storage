package repository

import (
	"context"
	"in_mem_storage/internal/domain/record"
	"in_mem_storage/internal/domain/ttl"
)

type ExpiryCandidate interface {
	Set(ctx context.Context, created record.CreatedAt, ec ttl.ExpiryCandidate) error
	Get(ctx context.Context, created record.CreatedAt) (ttl.ExpiryCandidate, error)
	Delete(ctx context.Context, created record.CreatedAt) error
}

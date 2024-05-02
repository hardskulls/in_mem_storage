package repository

import (
	"context"
	"in_mem_storage/internal/domain/log"
)

// LogRec is a repository for Log.
type LogRec[D any] interface {
	Save(ctx context.Context, lr log.Log[D]) error
}

package repository

import (
	"context"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/record"
)

type Record[D any] interface {
	Set(ctx context.Context, cmd command.Set[record.ID, D]) error
	Get(ctx context.Context, cmd command.Get[record.ID]) (record.Record[D], error)
	Update(ctx context.Context, cmd command.Update[record.ID, D]) error
	Delete(ctx context.Context, cmd command.Delete[record.ID]) error
}

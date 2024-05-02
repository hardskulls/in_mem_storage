package implementation

import (
	"context"
	"errors"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/record"
	"sync"
)

type RecordRepo[D any] struct {
	inner sync.Map
}

func (r *RecordRepo[D]) Set(
	ctx context.Context,
	cmd command.Set[record.ID, D],
) error {
	rec := record.New(cmd.Value(), cmd.Author())
	r.inner.Store(cmd.Key(), rec)
	return nil
}

func (r *RecordRepo[D]) Get(
	ctx context.Context,
	cmd command.Get[record.ID],
) (record.Record[D], error) {
	value, ok := r.inner.Load(cmd.Key())
	if !ok {
		return record.Record[D]{}, errors.New("key not found")
	}

	return value.(record.Record[D]), nil
}

func (r *RecordRepo[D]) Update(
	ctx context.Context,
	cmd command.Update[record.ID, D],
) error {
	value, ok := r.inner.Load(cmd.Key())
	if !ok {
		return errors.New("key not found")
	}

	rec := value.(record.Record[D])
	rec.Update(cmd.Value())

	r.inner.Store(cmd.Key(), rec)

	return nil
}

func (r *RecordRepo[D]) Delete(
	ctx context.Context,
	cmd command.Delete[record.ID],
) error {
	r.inner.Delete(cmd.Key())
	return nil
}

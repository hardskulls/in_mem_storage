package service

import (
	"context"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/data"
	"in_mem_storage/internal/domain/record"
	repository2 "in_mem_storage/internal/repository"
)

const (
	ok  = "OK"
	err = "ERR"
)

const (
	get = "GET"
	set = "SET"
	del = "DELETE"
	upd = "UPDATE"
)

type Command struct {
	record repository2.Record[string]
	ttl    repository2.ExpiryCandidate
}

func NewCommand(
	rec repository2.Record[string],
	ec repository2.ExpiryCandidate,
) Command {
	return Command{
		record: rec,
		ttl:    ec,
	}
}

func (cs *Command) Execute(
	ctx context.Context,
	cmd command.Command,
) (data.JSON, error) {
	switch c := cmd.(type) {
	case *command.Get[record.ID]:
		rec, err := cs.record.Get(ctx, *c)
		if err != nil {
			return data.JSON{}, err
		}
		b, _ := data.ToJson(command.WithPayload(ok, get, rec.Data()))
		return b, nil
	case *command.Set[record.ID, string]:
		err := cs.record.Set(ctx, *c)
		if err != nil {
			return data.JSON{}, err
		}
		b, _ := data.ToJson(command.NoPayload(ok, set))
		return b, nil
	case *command.Delete[record.ID]:
		err := cs.record.Delete(ctx, *c)
		if err != nil {
			return data.JSON{}, err
		}
		b, _ := data.ToJson(command.NoPayload(ok, del))
		return b, nil
	case *command.Update[record.ID, string]:
		err := cs.record.Update(ctx, *c)
		if err != nil {
			return data.JSON{}, err
		}
		b, _ := data.ToJson(command.NoPayload(ok, upd))
		return b, nil
	}

	return data.JSON{}, nil
}

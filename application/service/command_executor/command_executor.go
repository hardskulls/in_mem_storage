package services

import (
	obj "in_mem_storage/domain/transactions/commands/value_objects"
	repo "in_mem_storage/domain/transactions/record/repositories"
	rec "in_mem_storage/domain/transactions/record/value_objects"
)

type CommandService[K comparable, D, T, E any] struct {
	recordRepo repo.RecordRepo[K, D, T, E]
}

func WithRecordRepo[K comparable, D, T, E any](r repo.RecordRepo[K, D, T, E]) CommandService[K, D, T, E] {
	return CommandService[K, D, T, E]{recordRepo: r}
}

func (ce CommandService[K, D, T, E]) Get(cmd obj.GetCommand[K]) (rec.Record[D, T, E], error) {
	return ce.recordRepo.Get(cmd.Key())
}

func (ce CommandService[K, D, T, E]) Set(cmd obj.SetCommand[K, D, T, E]) error {
	return ce.recordRepo.Set(cmd.Key(), cmd.Value())
}

func (ce CommandService[K, D, T, E]) Delete(cmd obj.DeleteCommand[K]) error {
	return ce.recordRepo.Delete(cmd.Key())
}

func (ce CommandService[K, D, T, E]) Update(cmd obj.UpdateCommand[K, D, T, E]) error {
	return ce.recordRepo.Update(cmd.Key(), cmd.Value())
}

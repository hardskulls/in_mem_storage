package services

import (
	repo "in_mem_storage/application/service/command_executor/repository"
	cmd "in_mem_storage/domain/transaction/command/value_object"
	rec "in_mem_storage/domain/transaction/record/value_object"
)

type CommandService[K comparable, D, T, E any] struct {
	recordRepo repo.RecordRepo[K, D, T, E]
}

func WithRecordRepo[K comparable, D, T, E any](r repo.RecordRepo[K, D, T, E]) CommandService[K, D, T, E] {
	return CommandService[K, D, T, E]{recordRepo: r}
}

func (s CommandService[K, D, T, E]) Get(cmd cmd.GetCommand[K]) (rec.Record[D, T, E], error) {
	return s.recordRepo.Get(cmd.Key())
}

func (s CommandService[K, D, T, E]) Set(cmd cmd.SetCommand[K, D, T, E]) error {
	return s.recordRepo.Set(cmd.Key(), cmd.Value())
}

func (s CommandService[K, D, T, E]) Delete(cmd cmd.DeleteCommand[K]) error {
	return s.recordRepo.Delete(cmd.Key())
}

func (s CommandService[K, D, T, E]) Update(cmd cmd.UpdateCommand[K, D, T, E]) error {
	return s.recordRepo.Update(cmd.Key(), cmd.Value())
}

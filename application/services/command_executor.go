package services

import (
	// "in_mem_storage/internal/domain/errors"
	obj "in_mem_storage/internal/domain/transactions/commands/value_objects"
	// ts "in_mem_storage/internal/domain/transactions/abstractions"
	rec "in_mem_storage/internal/domain/transactions/record/value_objects"
	repo "in_mem_storage/internal/domain/transactions/record/repositories"
)

type CommandService[K comparable, D, T, E any] struct {
	recordRepo repo.RecordRepo[K, D, T, E]
}

func WithRecordRepo[K comparable, D, T, E any](r repo.RecordRepo[K, D, T, E]) CommandService[K, D, T, E] {
	return CommandService[K, D, T, E]{recordRepo: r}
}

func(ce CommandService[K, D, T, E]) Get(cmd obj.GetCommand[K]) (rec.Record[D, T, E], error) {
	return ce.recordRepo.Get(cmd.Key)
}

func(ce CommandService[K, D, T, E]) Set(cmd obj.SetCommand[K, D]) error {
	r := rec.NewRecord(cmd.Value, cmd.ExpiresIn)
	return ce.recordRepo.SetValue(cmd.Key, r)
}

func(ce CommandService[K, V]) Delete(cmd obj.GetCommand[K]) error {
	return ce.recordRepo.DeleteValue(cmd.Key)
}

// type CommandExecutor[OK any, K comparable] interface {
// 	Execute(with interface{ repo.RecordRepo[K] }) (OK, error)
// }

// func (s CommandService[OK, K]) ExecuteCmd(cmd interface {
// 	ts.CommandExecutor[rec.Record, K]
// }, r interface{ repo.RecordRepo[K] }) (rec.Record, error) {
// 	return cmd.Execute(r)
// }

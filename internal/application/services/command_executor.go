package services

import (
	// "in_mem_storage/internal/domain/errors"
	// ts "in_mem_storage/internal/domain/transactions/value_objects"
	repo "in_mem_storage/internal/domain/transactions/repositories"
)

type CommandExecutor[OK any, K comparable] interface {
	Execute(with interface{ repo.RecordRepo[K] }) (OK, error)
}

type CmdExecService[OK any, K comparable] struct{}

func (s CmdExecService[OK, K]) ExecuteCmd(cmd interface{ CommandExecutor[OK, K] }, r interface{ repo.RecordRepo[K] }) (OK, error) {
	return cmd.Execute(r)
}

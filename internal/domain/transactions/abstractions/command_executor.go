package abstractions

import (
	repo "in_mem_storage/internal/domain/transactions/repositories"
)

type CommandExecutor[OK any, K comparable] interface {
	Execute(with interface{ repo.RecordRepo[K] }) (OK, error)
}

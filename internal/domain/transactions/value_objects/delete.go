package value_objects

import (
	// rec "in_mem_storage/internal/domain/transactions/entities"
	// repo "in_mem_storage/internal/domain/transactions/repositories"
	repo "in_mem_storage/internal/domain/transactions/repositories"
)

type DeleteCommand[K comparable] struct {
	Key K
}

func (c *DeleteCommand[K]) Execute(r interface{ repo.RecordRepo[K] }) error {
	err := r.DeleteValue(c.Key)
	return err
}

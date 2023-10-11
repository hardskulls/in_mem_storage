package value_objects

import (
	rec "in_mem_storage/internal/domain/transactions/entities"
	// repo "in_mem_storage/internal/domain/transactions/repositories"
	repo "in_mem_storage/internal/domain/transactions/repositories"
)

type GetCommand[K comparable] struct {
	Key K
}

func (c *GetCommand[K]) Execute(r interface{ repo.RecordRepo[K] }) (rec.Record, error) {
	val, err := r.GetValue(c.Key)
	return val, err
}

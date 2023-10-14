package value_objects


import (
	// rec "in_mem_storage/internal/domain/transactions/entities"
	// repo "in_mem_storage/internal/domain/transactions/repositories"
	// repo "in_mem_storage/internal/domain/transactions/record/repositories"
)

type SetCommand[K comparable, V any] struct {
	Key       K
	Value     V
	ExpiresIn int
}

type GetCommand[K comparable] struct {
	Key K
}

type UpdateCommand0[K comparable, V any] struct {
	Key K
	Value V
}

type DeleteCommand[K comparable] struct {
	Key K
}


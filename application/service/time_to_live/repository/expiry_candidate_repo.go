package repository

import (
	ec "in_mem_storage/domain/time_to_live/value_object"
)

type ExpiryCandidateRepo[K, C comparable] interface {
	Set(key K, value ec.ExpiryCandidate[C]) error
	Get(key K) (ec.ExpiryCandidate[C], error)
	Delete(key K) error
}

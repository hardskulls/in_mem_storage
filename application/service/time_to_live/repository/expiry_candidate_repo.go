package repository

import (
	ec "in_mem_storage/domain/time_to_live/value_object"
)

type ExpiryCandidateRepo[K comparable] interface {
	Set(key K, value ec.ExpiryCandidate) error
	Get(key K) (ec.ExpiryCandidate, error)
	Delete(key K) error
}

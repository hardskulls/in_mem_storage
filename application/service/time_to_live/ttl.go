package services

import (
	repo "in_mem_storage/application/service/time_to_live/repository"
	ec "in_mem_storage/domain/time_to_live/value_object"
)

type TimeToLiveService[K comparable] struct {
	expiryCandidateRepo repo.ExpiryCandidateRepo[K]
}

func NewTimeToLiveService[K comparable](repo repo.ExpiryCandidateRepo[K]) TimeToLiveService[K] {
	return TimeToLiveService[K]{repo}
}

func (rl *TimeToLiveService[K]) Get(key K) (ec.ExpiryCandidate, error) {
	return rl.expiryCandidateRepo.Get(key)
}

func (rl *TimeToLiveService[K]) Set(key K, value ec.ExpiryCandidate) error {
	return rl.expiryCandidateRepo.Set(key, value)
}

func (rl *TimeToLiveService[K]) Delete(key K) error {
	return rl.expiryCandidateRepo.Delete(key)
}

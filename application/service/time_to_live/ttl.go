package services

import (
	repo "in_mem_storage/application/service/time_to_live/repository"
	ec "in_mem_storage/domain/time_to_live/value_object"
)

type TimeToLiveService[K, C comparable] struct {
	expiryCandidateRepo repo.ExpiryCandidateRepo[K, C]
}

func NewTimeToLiveService[K, C comparable](repo repo.ExpiryCandidateRepo[K, C]) TimeToLiveService[K, C] {
	return TimeToLiveService[K, C]{repo}
}

func (rl *TimeToLiveService[K, C]) Get(key K) (ec.ExpiryCandidate[C], error) {
	return rl.expiryCandidateRepo.Get(key)
}

func (rl *TimeToLiveService[K, C]) Set(key K, value ec.ExpiryCandidate[C]) error {
	return rl.expiryCandidateRepo.Set(key, value)
}

func (rl *TimeToLiveService[K, C]) Delete(key K) error {
	return rl.expiryCandidateRepo.Delete(key)
}

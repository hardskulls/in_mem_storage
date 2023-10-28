package services

import (
	repo "in_mem_storage/application/service/time_to_live/repository"
	ec "in_mem_storage/domain/time_to_live/value_object"
	"time"
)

type TimeToLiveService struct {
	expiryCandidateRepo repo.DefaultExpiryRecRepo
}

func New(repo repo.DefaultExpiryRecRepo) TimeToLiveService {
	return TimeToLiveService{repo}
}

func (rl *TimeToLiveService) Get(key time.Time) (ec.ExpiryRec, error) {
	return rl.expiryCandidateRepo.Get(key)
}

func (rl *TimeToLiveService) Set(key time.Time, value ec.ExpiryRec) error {
	return rl.expiryCandidateRepo.Set(key, value)
}

func (rl *TimeToLiveService) Delete(key time.Time) error {
	return rl.expiryCandidateRepo.Delete(key)
}

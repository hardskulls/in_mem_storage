package services

import (
	"fmt"
	cmdexec "in_mem_storage/internal/application/service/crud_cmd_executor/abstraction"
	repo "in_mem_storage/internal/application/service/crud_cmd_executor/repository"
	"in_mem_storage/internal/application/service/time_to_live/repository"
)

type CrudCommandService struct {
	recordRepo repo.DefaultRecordRepo
	ttlRepo    repository.DefaultExpiryRecRepo
}

func New(
	recRepo repo.DefaultRecordRepo,
	expiryRepo repository.DefaultExpiryRecRepo,
) CrudCommandService {
	return CrudCommandService{recRepo, expiryRepo}
}

func (s CrudCommandService) Execute(cmd cmdexec.DefaultCommandExecutor) fmt.Stringer {
	return cmd.Execute(s.recordRepo, s.ttlRepo)
}

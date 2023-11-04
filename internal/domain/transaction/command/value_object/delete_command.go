package value_objects

import (
	"fmt"
	rec "in_mem_storage/internal/application/service/crud_cmd_executor/repository"
	ttl "in_mem_storage/internal/application/service/time_to_live/repository"
)

type DeleteCommand struct {
	Key string
}

func (c DeleteCommand) Execute(
	recRepo rec.DefaultRecordRepo,
	_ ttl.DefaultExpiryRecRepo,
) fmt.Stringer {
	err := recRepo.Delete(c.Key)
	if err != nil {
		return UserNotification{Msg: err.Error()}
	}
	return UserNotification{Msg: "[DeleteCommand] executed successfully."}
}

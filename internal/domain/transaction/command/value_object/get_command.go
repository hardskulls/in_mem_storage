package value_objects

import (
	"fmt"
	rec "in_mem_storage/internal/application/service/crud_cmd_executor/repository"
	ttl "in_mem_storage/internal/application/service/time_to_live/repository"
)

type GetCommand struct {
	Key string
}

func (c GetCommand) Execute(
	recRepo rec.DefaultRecordRepo,
	_ ttl.DefaultExpiryRecRepo,
) fmt.Stringer {
	val, err := recRepo.Get(c.Key)
	if err != nil {
		return UserNotification{Msg: err.Error()}
	}
	return UserNotification{Msg: val.Data}
}

package value_objects

import (
	"fmt"
	recrepo "in_mem_storage/application/service/crud_cmd_executor/repository"
	ttl "in_mem_storage/application/service/time_to_live/repository"
	rec "in_mem_storage/domain/transaction/record/value_object"
)

type UpdateCommand struct {
	Key string
	Val string
}

func (c UpdateCommand) Execute(
	recRepo recrepo.DefaultRecordRepo,
	_ ttl.DefaultExpiryRecRepo,
) fmt.Stringer {
	err := recRepo.Update(c.Key, rec.Record{Data: c.Val})
	if err != nil {
		return UserNotification{Msg: err.Error()}
	}
	return UserNotification{Msg: "[UpdateCommand] executed successfully."}

}

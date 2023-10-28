package value_objects

import (
	"fmt"
	recrepo "in_mem_storage/application/service/crud_cmd_executor/repository"
	ttl "in_mem_storage/application/service/time_to_live/repository"
	expirycand "in_mem_storage/domain/time_to_live/value_object"
	rec "in_mem_storage/domain/transaction/record/value_object"
	"time"
)

type SetCommand struct {
	Key          string
	Val          string
	ExpiresAfter time.Duration
}

func (c SetCommand) Execute(
	recRepo recrepo.DefaultRecordRepo,
	ttlRepo ttl.DefaultExpiryRecRepo,
) fmt.Stringer {
	expiryDate := time.Now().Add(c.ExpiresAfter)
	err := ttlRepo.Set(expiryDate, expirycand.ExpiryRec{Record: c.Key})
	if err != nil {
		return UserNotification{Msg: fmt.Sprintf("[!! Error !!] : TTL not set, '%v'", err)}
	}

	err = recRepo.Set(c.Key, rec.Record{Data: c.Val})
	if err != nil {
		return UserNotification{Msg: fmt.Sprintf("[!! Error !!] : your data was not saved")}
	}

	return UserNotification{Msg: "[SetCommand] executed successfully."}

}

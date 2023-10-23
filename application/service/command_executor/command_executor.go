package services

import (
	"fmt"
	repo "in_mem_storage/application/service/command_executor/repository"
	"in_mem_storage/application/service/time_to_live/ttlrepo"
	expirycand "in_mem_storage/domain/time_to_live/value_object"
	c "in_mem_storage/domain/transaction/command/crud"
	commands "in_mem_storage/domain/transaction/command/value_object"
	rec "in_mem_storage/domain/transaction/record/value_object"
	"time"
)

type CrudCommandService[Command c.CrudCommand] struct {
	recordRepo repo.DefaultRecordRepo
	ttlRepo    ttlrepo.DefaultExpiryCandRepo
}

func New[C c.CrudCommand](
	recRepo repo.DefaultRecordRepo,
	expiryRepo ttlrepo.DefaultExpiryCandRepo,
) CrudCommandService[C] {
	return CrudCommandService[C]{recRepo, expiryRepo}
}

type userNotification struct {
	msg string
}

func (n userNotification) String() string {
	return n.msg
}

func (s CrudCommandService[C]) Execute(cmd C) fmt.Stringer {
	var a interface{} = cmd
	switch cmd := a.(type) {
	case commands.GetCommand:
		val, err := s.recordRepo.Get(cmd.Key)
		if err != nil {
			return userNotification{msg: err.Error()}
		}
		return userNotification{msg: val.Data}
	case commands.SetCommand:
		expiryDate := time.Now().Add(cmd.ExpiresAfter)
		err := s.ttlRepo.Set(expiryDate, expirycand.ExpiryCandidate{Candidate: cmd.Key})
		if err != nil {
			return userNotification{msg: fmt.Sprintf("[!! Error !!] : TTL not set, '%v'", err)}
		}

		err = s.recordRepo.Set(cmd.Key, rec.Record{Data: cmd.Val})
		if err != nil {
			return userNotification{msg: fmt.Sprintf("[!! Error !!] : your data was not saved")}
		}

		return userNotification{msg: "[SetCommand] executed successfully."}
	case commands.DeleteCommand:
		err := s.recordRepo.Delete(cmd.Key)
		if err != nil {
			return userNotification{msg: err.Error()}
		}
		return userNotification{msg: "[DeleteCommand] executed successfully."}
	case commands.UpdateCommand:
		err := s.recordRepo.Update(cmd.Key, rec.Record{Data: cmd.Val})
		if err != nil {
			return userNotification{msg: err.Error()}
		}
		return userNotification{msg: "[UpdateCommand] executed successfully."}
	default:
		return userNotification{msg: "[!! ERROR !!] : Command Executor hit default case."}
	}
}

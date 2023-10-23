package services

import (
	"fmt"
	repo "in_mem_storage/application/service/command_executor/repository"
	c "in_mem_storage/domain/transaction/command/crud"
	commands "in_mem_storage/domain/transaction/command/value_object"
	rec "in_mem_storage/domain/transaction/record/value_object"
)

//type CommandService[Key comparable, D fmt.Stringer] struct {
//	recordRepo repo.RecordRepo[Key, D]
//}
//
//func WithRecordRepo[Key comparable, D any](r repo.RecordRepo[Key, D]) CommandService[Key, D] {
//	return CommandService[Key, D]{recordRepo: r}
//}
//
//func (s CommandService[Key, D]) Get(cmd commands.GetCommand[Key]) (rec.Record[D], error) {
//	return s.recordRepo.Get(cmd.Key)
//}
//
//func (s CommandService[Key, D]) Set(cmd commands.SetCommand[Key, rec.Record[D]]) error {
//	return s.recordRepo.Set(cmd.Key, cmd.Val)
//}
//
//func (s CommandService[Key, D]) Delete(cmd commands.DeleteCommand[Key]) error {
//	return s.recordRepo.Delete(cmd.Key)
//}
//
//func (s CommandService[Key, D]) Update(cmd commands.UpdateCommand[Key, rec.Record[D]]) error {
//	return s.recordRepo.Update(cmd.Key, cmd.Val)
//}

type CrudCommandService[
	Command c.CrudCommand,
] struct {
	recordRepo repo.RecordRepo[string]
}

func New[
	C c.CrudCommand,
](with repo.RecordRepo[string]) CrudCommandService[C] {
	return CrudCommandService[C]{with}
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
		err := s.recordRepo.Set(cmd.Key, rec.Record{Data: cmd.Val})
		if err != nil {
			return userNotification{msg: err.Error()}
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

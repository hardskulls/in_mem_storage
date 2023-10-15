package command_executor

import (
	"fmt"
	errs "in_mem_storage/domain/error"
)

func DeleteError() errs.Error {
	msg := fmt.Errorf("[DeleteCommandError] : key did not exist").Error()
	levelsUp := 1
	return errs.New(msg, levelsUp)
}

func GetError() errs.Error {
	msg := fmt.Errorf("[GetCommandError] : couldn't acquire value").Error()
	levelsUp := 1
	return errs.New(msg, levelsUp)
}

func UpdateError() errs.Error {
	msg := fmt.Errorf("[UpdateCommandError] : key did not exist").Error()
	levelsUp := 1
	return errs.New(msg, levelsUp)
}

package repository

import (
	"fmt"
	errs "in_mem_storage/internal/domain/error/value_object"
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

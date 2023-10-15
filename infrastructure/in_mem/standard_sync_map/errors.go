package standard_sync_map

import (
	"fmt"
	"in_mem_storage/domain/errors"
)

func DeleteError() errors.Error {
	msg := fmt.Errorf("[DeleteCommandError] : key did not exist").Error()
	levelsUp := 1
	return errors.New(msg, levelsUp)
}

func GetError() errors.Error {
	msg := fmt.Errorf("[GetCommandError] : couldn't acquire value").Error()
	levelsUp := 1
	return errors.New(msg, levelsUp)
}

func UpdateError() errors.Error {
	msg := fmt.Errorf("[UpdateCommandError] : key did not exist").Error()
	levelsUp := 1
	return errors.New(msg, levelsUp)
}

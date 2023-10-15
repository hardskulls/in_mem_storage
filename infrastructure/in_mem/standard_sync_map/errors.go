package standard_sync_map

import (
	"fmt"
	"in_mem_storage/domain/errors"
)

func DeleteError() errors.Error {
	msg := fmt.Errorf("error: key did not exist").Error()
	levelsUp := 1
	return errors.New(msg, levelsUp)
}

func GetError() errors.Error {
	msg := fmt.Errorf("error: couldn't acquire value").Error()
	levelsUp := 1
	return errors.New(msg, levelsUp)
}

package abstraction

import (
	"in_mem_storage/internal/domain/transaction/command/value_object"
)

type CrudCommand interface {
	value_objects.GetCommand |
		value_objects.SetCommand |
		value_objects.DeleteCommand |
		value_objects.UpdateCommand
}

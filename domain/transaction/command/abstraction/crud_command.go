package abstraction

import c "in_mem_storage/domain/transaction/command/value_object"

type CrudCommand interface {
	c.GetCommand |
		c.SetCommand |
		c.DeleteCommand |
		c.UpdateCommand
}

package abstraction

import (
	cmdexec "in_mem_storage/internal/application/service/crud_cmd_executor/abstraction"
)

type CrudCommandProducer interface {
	ProduceCmd() (cmdexec.DefaultCommandExecutor, error)
}

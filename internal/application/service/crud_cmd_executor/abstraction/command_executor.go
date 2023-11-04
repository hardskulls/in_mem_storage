package abstraction

import (
	"fmt"
	rec "in_mem_storage/internal/application/service/crud_cmd_executor/repository"
	ttl "in_mem_storage/internal/application/service/time_to_live/repository"
	"time"
)

type DefaultCommandExecutor = CommandExecutor[time.Time, string, fmt.Stringer]

type CommandExecutor[T, K comparable, R any] interface {
	Execute(recRepo rec.RecordRepo[K], ttlRepo ttl.ExpiryRecRepo[T]) R
}

package port

import (
	"in_mem_storage/internal/domain/log/value_object/log_record"
)

type DefaultLogRecordPort = LogRecordPort[log_record.DefaultLogRecord]

type LogRecordPort[D any] interface {
	LogRecord(data D)
}

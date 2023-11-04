package logger

import (
	logrecrepo "in_mem_storage/internal/application/service/logger/port"
	logobj "in_mem_storage/internal/domain/log/value_object/log_record"
)

type Logger struct {
	logRecPort logrecrepo.DefaultLogRecordPort
}

func New(logRecPort logrecrepo.DefaultLogRecordPort) Logger {
	return Logger{logRecPort: logRecPort}
}

func (l Logger) Log(data logobj.DefaultLogRecord) {
	l.logRecPort.LogRecord(data)
}

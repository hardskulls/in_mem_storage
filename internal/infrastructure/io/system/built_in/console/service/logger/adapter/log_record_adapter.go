package console

import (
	logrecobj "in_mem_storage/internal/domain/log/value_object/log_record"
)

type LogRecordAdapter struct{}

func (c LogRecordAdapter) LogRecord(data logrecobj.DefaultLogRecord) {
	middleSeparator := "——————————————————————————————"
	separator := "==================================================================="

	logLvl := "LOG : " + string(data.LogLevel()) + "\n"
	event := "EVENT : " + string(data.Event()) + "\n"
	descr := "DESCRIPTION : " + data.Description() + "\n\n"
	logData := "  ->  " + data.Data() + "\n"
	created := "  ->  " + data.Created().String() + "\n"
	fileLine := "  ->  " + data.FileLine().String() + "\n\n"

	log := "\n" +
		separator + "\n\n" +
		logLvl +
		event +
		descr +
		middleSeparator + "\n\n" +
		logData +
		created +
		fileLine +
		separator + "\n\n\n"

	println(log)
}

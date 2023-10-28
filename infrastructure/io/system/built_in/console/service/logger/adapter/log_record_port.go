package console

import (
	logrecobj "in_mem_storage/domain/log/value_object/log_record"
)

type LogRecordAdapter struct{}

func (c LogRecordAdapter) LogRecord(data logrecobj.DefaultLogRecord) {
	middleSeparator := "——————————————————————————————"
	separator := "==================================================================="

	log := "\n" +
		separator + "\n\n" +
		"LOG : " + data.LogLevel + "\n" +
		"EVENT : " + data.Event + "\n" +
		"DESCRIPTION : " + data.Description + "\n\n" +
		middleSeparator + "\n\n" +
		"  ->  " + data.Data.Error() + "\n" +
		"  ->  " + data.EventOccurredAt.String() + "\n" +
		"  ->  " + data.FileLine.String() + "\n\n" +
		separator + "\n\n\n"

	println(log)
}

package value_object

type LogLvl string

const (
	Error   LogLvl = "error"
	Warning LogLvl = "warning"
	Info    LogLvl = "info"
	Debug   LogLvl = "debug"
	Trace   LogLvl = "trace"
)

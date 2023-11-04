package value_object

import (
	"fmt"
	fl "in_mem_storage/internal/domain/log/value_object/file_line"
	"runtime"
)

type Error struct {
	msg  string
	file string
	line int
}

func New(msg string, levelsFromHere int) Error {
	_, file, line, _ := runtime.Caller(levelsFromHere + 1)
	return Error{msg, file, line}
}

func FromError(err error, levelsFromHere int) Error {
	_, file, line, _ := runtime.Caller(levelsFromHere + 1)
	return Error{err.Error(), file, line}
}

func (e Error) Msg() string {
	return e.msg
}

func (e Error) FileLine() fl.FileLine {
	return fl.FileLine{File: e.file, Line: e.line}
}

func (e Error) OccurredAt() string {
	return fmt.Sprintf("%v:%v", e.file, e.line)
}

func (e Error) Error() string {
	return fmt.Sprintf("msg: '%v', file and line: %v", e.Msg(), e.OccurredAt())
}

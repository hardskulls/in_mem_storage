package errors

import (
	"fmt"
	"runtime"
)

type Error struct {
	msg  string
	file string
	line int
}

func New(msg string, skip int) Error {
	_, file, line, _ := runtime.Caller(skip + 1)
	return Error{msg: msg, file: file, line: line}
}

func (e Error) Msg() string {
	return e.msg
}

func (e Error) OccurredAt() string {
	return fmt.Sprintf("%v:%v", e.file, e.line)
}

func (e Error) Error() string {
	return fmt.Sprintf("msg: '%v', file and line: %v", e.Msg(), e.OccurredAt())
}

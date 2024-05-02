package locerr

import (
	"fmt"
	"in_mem_storage/internal/domain/log"
	"runtime"
)

type Message = string
type OccurredAt = string

type Result[T, E any] struct {
	Ok  T
	Err E
}

// Error is a local error value object.
type Error struct {
	msg  Message
	file log.File
	line log.Line
}

// New returns a local error instance.
func New(msg string, up log.Where) Error {
	_, file, line, _ := runtime.Caller(int(up + 1))
	return Error{msg, file, line}
}

// FromError creates anew error from an existing one.
func FromError(err error, up log.Where) Error {
	_, file, line, _ := runtime.Caller(int(up + 1))
	return Error{err.Error(), file, line}
}

func (e Error) Msg() Message {
	return e.msg
}

func (e Error) CodeLoc() log.CodeLoc {
	return log.NewCodeLoc(log.InOuterFn)
}

func (e Error) OccurredAt() OccurredAt {
	return fmt.Sprintf("%v:%v", e.file, e.line)
}

func (e Error) Error() string {
	return fmt.Sprintf("msg: '%v', location: %v", e.Msg(), e.OccurredAt())
}

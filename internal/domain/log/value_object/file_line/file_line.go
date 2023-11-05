package file_line

import (
	"fmt"
	stfrup "in_mem_storage/internal/domain/log/value_object/stack_frames_up"
	"runtime"
)

type FileLine struct {
	file string
	line int
}

func New(levelsFromHere stfrup.StackFramesUp) FileLine {
	_, file, line, _ := runtime.Caller(int(levelsFromHere) + 1)
	return FileLine{file, line}
}

func (fl FileLine) String() string {
	return fmt.Sprintf("%v:%v", fl.file, fl.line)
}

func (fl FileLine) File() string {
	return fl.file
}

func (fl FileLine) Line() int {
	return fl.line
}

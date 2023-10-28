package file_line

import (
	"fmt"
	"runtime"
)

type FileLine struct {
	File string
	Line int
}

func New(levelsFromHere int) FileLine {
	_, file, line, _ := runtime.Caller(levelsFromHere + 1)
	return FileLine{file, line}
}

func (fl FileLine) String() string {
	return fmt.Sprintf("%v:%v", fl.File, fl.Line)
}

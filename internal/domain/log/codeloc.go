package log

import (
	"fmt"
	"runtime"
)

type File = string
type Line = int

// CodeLoc is a value object representing a location in the codebase.
type CodeLoc struct {
	file File
	line Line
}

// NewCodeLoc returns anew CodeLoc.
func NewCodeLoc(levelsFromHere Where) CodeLoc {
	_, file, line, _ := runtime.Caller(int(levelsFromHere) + 1)
	return CodeLoc{file, line}
}

func (fl CodeLoc) String() string {
	return fmt.Sprintf("%v:%v", fl.file, fl.line)
}

func (fl CodeLoc) File() File {
	return fl.file
}

func (fl CodeLoc) Line() Line {
	return fl.line
}

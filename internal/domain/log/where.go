package log

// Where is a value object that defines where to log:
// right here or 1 level up from here.
type Where int

const (
	Here Where = iota
	InOuterFn
)

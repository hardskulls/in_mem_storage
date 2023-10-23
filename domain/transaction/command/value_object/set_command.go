package value_objects

import "time"

type SetCommand struct {
	Key          string
	Val          string
	ExpiresAfter time.Duration
}

package value_objects

import "time"

type Record struct {
	Data    string
	Author  string
	Created time.Time
}

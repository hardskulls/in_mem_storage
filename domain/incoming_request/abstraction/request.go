package abstraction

import "time"

type Request[Body any] interface {
	Date() time.Time
	From() string
	Body() Body
}

package value_objects

// import "time"

type RateLimit[U, L any] struct {
	lastUse U
	limit   L
}

func NewRateLimit[U, L any](lastUse U, limit L) RateLimit[U, L] {
	return RateLimit[U, L]{lastUse: lastUse, limit: limit}
}

func (rl RateLimit[U, _]) LastUse() U {
	return rl.lastUse
}

func (rl RateLimit[_, L]) Limit() L {
	return rl.limit
}

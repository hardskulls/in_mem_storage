package abstractions

// "time"

type TokenStorage[C comparable] interface {
	AvailableTokens(consumer C, spend int) int
}

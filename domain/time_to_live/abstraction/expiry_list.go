package abstractions

type ExpiryList[Item any] interface {
	Add(item Item) error
	Remove(item Item) error
}

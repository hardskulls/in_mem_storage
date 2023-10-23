package crud

type CrudCommandProducer[C CrudCommand] interface {
	ProduceCmd() (C, error)
}

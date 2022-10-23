package data

type GeneratorDB interface {
	New() GeneratorDB

	Tasks() TasksQ
	KeyValue() KeyValueQ

	Transaction(func() error) error
}

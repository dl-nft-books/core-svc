package data

type GeneratorDB interface {
	New() GeneratorDB

	Tasks() TasksQ
	KeyValue() KeyValueQ
	Tokens() TokensQ

	Transaction(func() error) error
}

package data

type GeneratorDB interface {
	New() GeneratorDB

	Tasks() TasksQ
	KeyValue() KeyValueQ
	Tokens() TokensQ
	Attempts() AttemptsQ

	Transaction(func() error) error
}

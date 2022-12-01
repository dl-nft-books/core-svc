package data

type DB interface {
	New() DB

	Tasks() TasksQ
	KeyValue() KeyValueQ
	Tokens() TokensQ

	Transaction(func() error) error
}

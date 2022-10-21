package data

type DB interface {
	New() DB

	Tasks() TasksQ
	KeyValue() KeyValueQ
	Books() BookQ

	Transaction(func() error) error
}

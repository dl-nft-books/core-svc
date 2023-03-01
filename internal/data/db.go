package data

type DB interface {
	New() DB

	Tasks() TasksQ
	KeyValue() KeyValueQ
	Tokens() TokensQ
	Promocodes() PromocodesQ

	Transaction(func() error) error
}

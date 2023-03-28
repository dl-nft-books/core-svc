package data

type DB interface {
	New() DB

	Tasks() TasksQ
	KeyValue() KeyValueQ
	Promocodes() PromocodesQ
	NftRequests() NftRequestsQ
	PromocodesBooks() PromocodesBooksQ

	Transaction(func() error) error
}

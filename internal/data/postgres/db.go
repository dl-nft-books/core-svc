package postgres

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type db struct {
	raw *pgdb.DB
}

func NewDB(rawDB *pgdb.DB) data.DB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.DB {
	return NewDB(db.raw.Clone())
}

func (db *db) Tasks() data.TasksQ {
	return NewTasksQ(db.raw)
}

func (db *db) KeyValue() data.KeyValueQ {
	return NewKeyValueQ(db.raw)
}

func (db *db) Promocodes() data.PromocodesQ {
	return NewPromocodesQ(db.raw)
}

func (db *db) PromocodesBooks() data.PromocodesBooksQ {
	return NewPromocodesBooksQ(db.raw)
}

func (db *db) NftRequests() data.NftRequestsQ {
	return NewNftRequestsQ(db.raw)
}

func (db *db) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}

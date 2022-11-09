package postgres

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
)

type db struct {
	raw *pgdb.DB
}

func NewGeneratorDB(rawDB *pgdb.DB) data.GeneratorDB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.GeneratorDB {
	return NewGeneratorDB(db.raw.Clone())
}

func (db *db) Tasks() data.TasksQ {
	return NewTasksQ(db.raw)
}

func (db *db) KeyValue() data.KeyValueQ {
	return NewKeyValueQ(db.raw)
}

func (db *db) Tokens() data.TokensQ {
	return NewTokensQ(db.raw)
}

func (db *db) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}

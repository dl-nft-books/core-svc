package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
)

const booksTableName = "book"
const priceColumnName = "price"

func NewBooksQ(db *pgdb.DB) data.BookQ {
	return &BooksQ{
		db: db.Clone(),
		sql: squirrel.
			Select("*").
			From(fmt.Sprintf("%s as b", booksTableName)),
	}
}

type BooksQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func (b *BooksQ) New() data.BookQ {
	return NewBooksQ(b.db)
}

func (b *BooksQ) Get() (*data.Book, error) {
	var result data.Book

	err := b.db.Get(&result, b.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (b *BooksQ) FilterByID(id int64) data.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.id": id})
	return b
}

func (b *BooksQ) FilterActual() data.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.deleted": "f"})
	return b
}

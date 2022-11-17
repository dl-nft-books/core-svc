package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/external"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	booksTableName = "book"
	booksPrice     = "price"
)

func NewBooksQ(db *pgdb.DB) external.BookQ {
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

func (b *BooksQ) New() external.BookQ {
	return NewBooksQ(b.db)
}

func (b *BooksQ) Get() (*external.Book, error) {
	var result external.Book

	err := b.db.Get(&result, b.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (b *BooksQ) FilterByID(id int64) external.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.id": id})
	return b
}

func (b *BooksQ) FilterActual() external.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.deleted": "f"})
	return b
}

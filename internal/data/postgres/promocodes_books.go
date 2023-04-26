package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/core-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	promocodesBooksTable       = "promocodes_books"
	promocodesBooksPromocodeId = "promocode_id"
	promocodesBooksBookId      = "book_id"
)

type promocodesBooksQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewPromocodesBooksQ(database *pgdb.DB) data.PromocodesBooksQ {
	return &promocodesBooksQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", promocodesBooksTable)).From(promocodesBooksTable),
	}
}

func (q *promocodesBooksQ) New() data.PromocodesBooksQ {
	return NewPromocodesBooksQ(q.database.Clone())
}

func (q *promocodesBooksQ) Insert(promocodeId int64, bookIds ...int64) error {
	statement := squirrel.Insert(promocodesBooksTable).Columns(promocodesBooksPromocodeId, promocodesBooksBookId)
	for _, bookId := range bookIds {
		statement = statement.Values(promocodeId, bookId)
	}
	err := q.database.Exec(statement)
	return err
}

func (q *promocodesBooksQ) FilterByPromocodeId(id ...int64) data.PromocodesBooksQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesBooksPromocodeId: id})
	return q
}
func (q *promocodesBooksQ) FilterByBookId(id ...int64) data.PromocodesBooksQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesBooksBookId: id})
	return q
}

func (q *promocodesBooksQ) DeleteBooks(promocodeId int64, bookIds ...int64) error {
	return q.database.Exec(squirrel.Delete(promocodesBooksTable).
		Where(squirrel.Eq{promocodesBooksPromocodeId: promocodeId}).
		Where(squirrel.Eq{promocodesBooksBookId: bookIds}))
}

func (q *promocodesBooksQ) SelectBooks() (promocodes []int64, err error) {
	err = q.database.Select(&promocodes, q.selector.Columns(promocodesBooksBookId))
	return
}

func (q *promocodesBooksQ) GetBooks() (*int64, error) {
	var promocode int64
	err := q.database.Get(&promocode, q.selector.Columns(promocodesBooksBookId))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &promocode, err
}
func (q *promocodesBooksQ) SelectPromocodes() (promocodes []int64, err error) {
	err = q.database.Select(&promocodes, q.selector.Columns(promocodesBooksPromocodeId))
	return
}

func (q *promocodesBooksQ) GetPromocodes() (*int64, error) {
	var promocode int64
	err := q.database.Get(&promocode, q.selector.Columns(promocodesBooksPromocodeId))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &promocode, err
}

func (q *promocodesBooksQ) UpdateBooks(promocodeId int64, bookIds ...int64) error {
	insertStatement := squirrel.Insert(promocodesBooksTable).Columns(promocodesBooksPromocodeId, promocodesBooksBookId)
	for _, bookId := range bookIds {
		insertStatement = insertStatement.Values(promocodeId, bookId)
	}

	err := q.database.Exec(insertStatement.Suffix("ON CONFLICT DO NOTHING"))
	if err != nil {
		return err
	}

	return q.database.Exec(squirrel.Delete("promocodes_books").
		Where(squirrel.Eq{"promocode_id": promocodeId}).
		Where(squirrel.NotEq{"book_id": bookIds}))
}

func (q *promocodesBooksQ) UpdatePromocodes(bookId int64, promocodeIds ...int64) error {
	insertStatement := squirrel.Insert(promocodesBooksTable).Columns(promocodesBooksPromocodeId, promocodesBooksBookId)
	for _, promocodeId := range promocodeIds {
		insertStatement = insertStatement.Values(promocodeId, bookId)
	}

	err := q.database.Exec(insertStatement.Suffix("ON CONFLICT DO NOTHING"))
	if err != nil {
		return err
	}

	return q.database.Exec(squirrel.Delete("promocodes_books").
		Where(squirrel.Eq{"promocode_id": bookId}).
		Where(squirrel.NotEq{"book_id": promocodeIds}))
}

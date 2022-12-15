package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	promocodesTable = "promocodes"

	promocodesId             = "id"
	promocodesPromocode      = "promocode"
	promocodesDiscount       = "discount"
	promocodesInitialUsages  = "initial_usages"
	promocodesLeftUsages     = "left_usages"
	promocodesExpirationDate = "expiration_date"
	promocodesState          = "state"
)

type promocodesQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	deleter  squirrel.DeleteBuilder
}

func NewPromocodesQ(database *pgdb.DB) data.PromocodesQ {
	return &promocodesQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", promocodesTable)).From(promocodesTable),
		updater:  squirrel.Update(promocodesTable).Suffix("RETURNING *"),
		deleter:  squirrel.Delete(promocodesTable),
	}
}

func (q *promocodesQ) New() data.PromocodesQ {
	return NewPromocodesQ(q.database.Clone())
}

func (q *promocodesQ) Page(page pgdb.OffsetPageParams) data.PromocodesQ {
	q.selector = page.ApplyTo(q.selector, "id")
	return q
}

func (q *promocodesQ) Sort(sort pgdb.Sorts) data.PromocodesQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *promocodesQ) FilterById(id ...int64) data.PromocodesQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesId: id})
	q.deleter = q.deleter.Where(squirrel.Eq{promocodesId: id})
	return q
}

func (q *promocodesQ) FilterByPromocode(promocode ...string) data.PromocodesQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesPromocode: promocode})
	q.deleter = q.deleter.Where(squirrel.Eq{promocodesPromocode: promocode})
	return q
}

func (q *promocodesQ) FilterByState(state ...resources.PromocodeState) data.PromocodesQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesState: state})
	q.deleter = q.deleter.Where(squirrel.Eq{promocodesState: state})

	return q
}

func (q *promocodesQ) Select() (promocodes []data.Promocode, err error) {
	err = q.database.Select(&promocodes, q.selector)
	return
}

func (q *promocodesQ) Get() (*data.Promocode, error) {
	var promocode data.Promocode
	err := q.database.Get(&promocode, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &promocode, err
}

func (q *promocodesQ) Delete() error {
	var promocode data.Promocode
	err := q.database.Get(&promocode, q.deleter)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func (q *promocodesQ) Insert(promocode data.Promocode) (data.Promocode, error) {
	statement := squirrel.Insert(promocodesTable).
		Suffix("returning *").
		SetMap(structs.Map(&promocode))

	err := q.database.Get(&promocode, statement)
	return promocode, err
}

func (q *promocodesQ) Transaction(fn func(q data.PromocodesQ) error) (err error) {
	return q.database.Transaction(func() error {
		return fn(q)
	})
}

func (q *promocodesQ) UpdateState(newState resources.PromocodeState) data.PromocodesQ {
	q.updater = q.updater.Set(promocodesState, newState)
	return q
}

func (q *promocodesQ) UpdateDiscount(newDiscount float64) data.PromocodesQ {
	q.updater = q.updater.Set(promocodesDiscount, newDiscount)
	return q
}

func (q *promocodesQ) UpdateInitialUsages(newInitialUsages int64) data.PromocodesQ {
	q.updater = q.updater.Set(promocodesInitialUsages, newInitialUsages)
	return q
}

func (q *promocodesQ) UpdateLeftUsages(newLeftUsages int64) data.PromocodesQ {
	q.updater = q.updater.Set(promocodesLeftUsages, newLeftUsages)
	return q
}

func (q *promocodesQ) UpdateExpirationDate(newExpirationDate time.Time) data.PromocodesQ {
	q.updater = q.updater.Set(promocodesExpirationDate, newExpirationDate)
	return q
}

func (q *promocodesQ) Update(id int64) (*data.Promocode, error) {
	var promocode data.Promocode
	err := q.database.Get(&promocode, q.updater.Where(squirrel.Eq{promocodesId: id}))
	return &promocode, err
}

func (q *promocodesQ) UpdateWhereExpired() error {
	err := q.database.Exec(q.updater.
		Where(squirrel.LtOrEq{promocodesExpirationDate: time.Now()}).
		Where(squirrel.Eq{promocodesState: resources.PromocodeActive}))
	return err
}

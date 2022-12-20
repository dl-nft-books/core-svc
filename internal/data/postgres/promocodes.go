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
}

func NewPromocodesQ(database *pgdb.DB) data.PromocodesQ {
	return &promocodesQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", promocodesTable)).From(promocodesTable),
		updater:  squirrel.Update(promocodesTable).Suffix("RETURNING *"),
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
	return q
}

func (q *promocodesQ) FilterByPromocode(promocode ...string) data.PromocodesQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesPromocode: promocode})
	return q
}

func (q *promocodesQ) FilterByState(state ...resources.PromocodeState) data.PromocodesQ {
	q.selector = q.selector.Where(squirrel.Eq{promocodesState: state})

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

func (q *promocodesQ) DeleteByID(id int64) error {
	return q.database.Exec(squirrel.Delete(promocodesTable).
		Where(squirrel.Eq{promocodesId: id}))
}

func (q *promocodesQ) Insert(promocode data.Promocode) (int64, error) {
	var id int64
	statement := squirrel.Insert(promocodesTable).
		Suffix("returning id").
		SetMap(structs.Map(&promocode))

	err := q.database.Get(&id, statement)
	return id, err
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

func (q *promocodesQ) Update(id int64) error {
	return q.database.Exec(q.updater.Where(squirrel.Eq{promocodesId: id}))
}

func (q *promocodesQ) UpdateWhereExpired() error {
	return q.database.Exec(q.updater.
		Where(squirrel.LtOrEq{promocodesExpirationDate: time.Now()}).
		Where(squirrel.Eq{promocodesState: resources.PromocodeActive}))
}

func (q *promocodesQ) UpdateWhereFullyUsed() error {
	return q.database.Exec(q.updater.
		Where(squirrel.LtOrEq{promocodesLeftUsages: 0}).
		Where(squirrel.Eq{promocodesState: resources.PromocodeActive}))
}

func (q *promocodesQ) UpdateWhereActive() error {
	return q.database.Exec(q.updater.
		Where(squirrel.Gt{promocodesExpirationDate: time.Now()}).
		Where(squirrel.Gt{promocodesLeftUsages: 0}).
		Where(squirrel.NotEq{promocodesState: resources.PromocodeActive}))
}

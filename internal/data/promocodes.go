package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"time"
)

type Promocode struct {
	Id             int64                    `db:"id" structs:"-" json:"-"`
	Promocode      string                   `db:"promocode" structs:"promocode"`
	Discount       float64                  `db:"discount" structs:"discount"`
	InitialUsages  int64                    `db:"initial_usages" structs:"initial_usages" json:"initial_usages"`
	LeftUsages     int64                    `db:"left_usages" structs:"left_usages" json:"left_usages"`
	ExpirationDate time.Time                `db:"expiration_date" structs:"expiration_date" json:"expiration_date"`
	State          resources.PromocodeState `db:"state" structs:"state" json:"state"`
}

type PromocodesQ interface {
	New() PromocodesQ

	Get() (*Promocode, error)
	Select() ([]Promocode, error)
	Delete() error

	Sort(sort pgdb.Sorts) PromocodesQ
	Page(page pgdb.OffsetPageParams) PromocodesQ

	Insert(promocode Promocode) (Promocode, error)
	Transaction(fn func(q PromocodesQ) error) error
	FilterByState(status ...resources.PromocodeState) PromocodesQ
	FilterById(id ...int64) PromocodesQ
	FilterByPromocode(promocode ...string) PromocodesQ

	UpdateState(newState resources.PromocodeState) PromocodesQ
	UpdateDiscount(newDiscount float64) PromocodesQ
	UpdateInitialUsages(newInitialUsages int64) PromocodesQ
	UpdateLeftUsages(newLeftUsages int64) PromocodesQ
	UpdateExpirationDate(newExpirationDate time.Time) PromocodesQ
	Update(id int64) (*Promocode, error)
	UpdateWhereExpired() error
}

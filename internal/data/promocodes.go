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
	Usages     int64                    `db:"usages" structs:"usages" json:"usages"`
	ExpirationDate time.Time                `db:"expiration_date" structs:"expiration_date" json:"expiration_date"`
	State          resources.PromocodeState `db:"state" structs:"state" json:"state"`
}

func (promocode *Promocode) Resource() resources.Promocode {
	return resources.Promocode{
		Key: resources.NewKeyInt64(promocode.Id, resources.PROMOCODE),
		Attributes: resources.PromocodeAttributes{
			Id:             promocode.Id,
			Promocode:      promocode.Promocode,
			Discount:       promocode.Discount,
			InitialUsages:  promocode.InitialUsages,
			Usages:     promocode.Usages,
			ExpirationDate: promocode.ExpirationDate,
			State:          promocode.State,
		},
	}
}

type PromocodesQ interface {
	New() PromocodesQ

	Get() (*Promocode, error)
	Select() ([]Promocode, error)
	DeleteByID(id int64) error

	Sort(sort pgdb.Sorts) PromocodesQ
	Page(page pgdb.OffsetPageParams) PromocodesQ

	Insert(promocode Promocode) (int64, error)
	Transaction(fn func(q PromocodesQ) error) error
	FilterByState(state ...resources.PromocodeState) PromocodesQ
	FilterById(id ...int64) PromocodesQ
	FilterByPromocode(promocode ...string) PromocodesQ

	UpdateState(newState resources.PromocodeState) PromocodesQ
	UpdateDiscount(newDiscount float64) PromocodesQ
	UpdateInitialUsages(newInitialUsages int64) PromocodesQ
	UpdateUsages(newUsages int64) PromocodesQ
	UpdateExpirationDate(newExpirationDate time.Time) PromocodesQ
	Update(id int64) error
	UpdateWhereExpired() error
	UpdateWhereFullyUsed() error
	UpdateWhereActive() error
}

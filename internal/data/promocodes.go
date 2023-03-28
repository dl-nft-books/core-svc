package data

import (
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type Promocode struct {
	Id             int64                    `db:"id" structs:"-"`
	Promocode      string                   `db:"promocode" structs:"promocode"`
	Discount       float64                  `db:"discount" structs:"discount"`
	InitialUsages  int64                    `db:"initial_usages" structs:"initial_usages"`
	Usages         int64                    `db:"usages" structs:"usages"`
	ExpirationDate time.Time                `db:"expiration_date" structs:"expiration_date"`
	State          resources.PromocodeState `db:"state" structs:"state"`
}

func (promocode *Promocode) Resource() resources.Promocode {
	return resources.Promocode{
		Key: resources.NewKeyInt64(promocode.Id, resources.PROMOCODE),
		Attributes: resources.PromocodeAttributes{
			Id:             promocode.Id,
			Promocode:      promocode.Promocode,
			Discount:       promocode.Discount,
			InitialUsages:  promocode.InitialUsages,
			Usages:         promocode.Usages,
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
	FilterUpdateById(id ...int64) PromocodesQ
	FilterByPromocode(promocode ...string) PromocodesQ

	UpdateState(newState resources.PromocodeState) PromocodesQ
	UpdateDiscount(newDiscount float64) PromocodesQ
	UpdateInitialUsages(newInitialUsages int64) PromocodesQ
	UpdateUsages(newUsages int64) PromocodesQ
	UpdateExpirationDate(newExpirationDate time.Time) PromocodesQ
	UpdatePromocode(promocode string) PromocodesQ

	Update() error
	FilterExpired() PromocodesQ
	FilterFullyUsed() PromocodesQ
	FilterActive() PromocodesQ
}

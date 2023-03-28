package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	nftRequestsTable = "nft_requests"

	nftRequestsId                = "id"
	nftRequestsPayerAddress      = "payer_address"
	nftRequestsCollectionAddress = "collection_address"
	nftRequestsFloorPrice        = "floor_price"
	nftRequestsNftId             = "nft_id"
	nftRequestsChainId           = "chain_id"
	nftRequestsBookId            = "book_id"
	nftRequestsCreatedAt         = "created_at"
	nftRequestsStatus            = "status"
)

type nftRequestsQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewNftRequestsQ(database *pgdb.DB) data.NftRequestsQ {
	return &nftRequestsQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", nftRequestsTable)).From(nftRequestsTable),
		updater:  squirrel.Update(nftRequestsTable).Suffix("RETURNING *"),
	}
}

func (q *nftRequestsQ) New() data.NftRequestsQ {
	return NewNftRequestsQ(q.database.Clone())
}

func (q *nftRequestsQ) Page(page pgdb.OffsetPageParams) data.NftRequestsQ {
	q.selector = page.ApplyTo(q.selector, "id")
	return q
}

func (q *nftRequestsQ) Sort(sort pgdb.Sorts) data.NftRequestsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *nftRequestsQ) FilterById(id ...int64) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsId: id})
	return q
}

func (q *nftRequestsQ) FilterByChainId(id ...int64) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsChainId: id})
	return q
}

func (q *nftRequestsQ) FilterUpdateById(id ...int64) data.NftRequestsQ {
	q.updater = q.updater.Where(squirrel.Eq{nftRequestsId: id})
	return q
}

func (q *nftRequestsQ) FilterByBookId(id ...int64) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsBookId: id})
	return q
}

func (q *nftRequestsQ) FilterByNftId(id ...int64) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsNftId: id})
	return q
}

func (q *nftRequestsQ) FilterByCollectionAddress(address ...string) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsCollectionAddress: address})
	return q
}

func (q *nftRequestsQ) FilterByPayerAddress(address ...string) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsPayerAddress: address})
	return q
}

func (q *nftRequestsQ) FilterByStatus(status ...resources.NftRequestStatus) data.NftRequestsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftRequestsStatus: status})

	return q
}

func (q *nftRequestsQ) Select() (nftRequests []data.NftRequest, err error) {
	err = q.database.Select(&nftRequests, q.selector)
	return
}

func (q *nftRequestsQ) Get() (*data.NftRequest, error) {
	var nftRequest data.NftRequest
	err := q.database.Get(&nftRequest, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &nftRequest, err
}

func (q *nftRequestsQ) DeleteByID(id int64) error {
	return q.database.Exec(squirrel.Delete(nftRequestsTable).
		Where(squirrel.Eq{nftRequestsId: id}))
}

func (q *nftRequestsQ) Insert(nftRequest data.NftRequest) (int64, error) {
	var id int64
	statement := squirrel.Insert(nftRequestsTable).
		Suffix("returning id").
		SetMap(structs.Map(&nftRequest))

	err := q.database.Get(&id, statement)
	return id, err
}

func (q *nftRequestsQ) Transaction(fn func(q data.NftRequestsQ) error) (err error) {
	return q.database.Transaction(func() error {
		return fn(q)
	})
}

func (q *nftRequestsQ) UpdateStatus(newState resources.NftRequestStatus) data.NftRequestsQ {
	q.updater = q.updater.Set(nftRequestsStatus, newState)
	return q
}

func (q *nftRequestsQ) Update() error {
	return q.database.Exec(q.updater)
}

package data

import (
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type NftRequest struct {
	Id                   int64                      `db:"id" structs:"-"`
	Requester            string                     `db:"requester" structs:"requester"`
	MarketplaceRequestId int64                      `db:"marketplace_request_id" structs:"marketplace_request_id"`
	NftAddress           string                     `db:"nft_address" structs:"nft_address"`
	NftId                int64                      `db:"nft_id" structs:"nft_id"`
	ChainId              int64                      `db:"chain_id" structs:"chain_id"`
	BookId               int64                      `db:"book_id" structs:"book_id"`
	Status               resources.NftRequestStatus `db:"status" structs:"status"`
	CreatedAt            time.Time                  `db:"created_at" structs:"created_at"`
	LastUpdatedAt        time.Time                  `db:"last_updated_at" structs:"last_updated_at"`
}

func (nftRequest *NftRequest) Resource() resources.NftRequest {
	book := resources.NewKeyInt64(nftRequest.BookId, resources.BOOKS)
	return resources.NftRequest{
		Key: resources.NewKeyInt64(nftRequest.Id, resources.PROMOCODE),
		Attributes: resources.NftRequestAttributes{
			ChainId:              nftRequest.ChainId,
			MarketplaceRequestId: nftRequest.MarketplaceRequestId,
			NftAddress:           nftRequest.NftAddress,
			NftId:                nftRequest.NftId,
			Requester:            nftRequest.Requester,
			Status:               nftRequest.Status,
			CreatedAt:            nftRequest.CreatedAt,
			LastUpdatedAt:        nftRequest.LastUpdatedAt,
		},
		Relationships: resources.NftRequestRelationships{
			Book: resources.Relation{
				Data: &book,
			},
		},
	}
}

type NftRequestsQ interface {
	New() NftRequestsQ

	Get() (*NftRequest, error)
	Select() ([]NftRequest, error)
	DeleteByID(id int64) error

	Sort(sort pgdb.Sorts) NftRequestsQ
	Page(page pgdb.OffsetPageParams) NftRequestsQ

	Insert(nftRequest NftRequest) (int64, error)
	Transaction(fn func(q NftRequestsQ) error) error
	FilterByStatus(status ...resources.NftRequestStatus) NftRequestsQ
	FilterById(id ...int64) NftRequestsQ
	FilterByChainId(id ...int64) NftRequestsQ
	FilterByBookId(id ...int64) NftRequestsQ
	FilterByNftId(id ...int64) NftRequestsQ
	FilterByNftAddress(address ...string) NftRequestsQ
	FilterByRequester(address ...string) NftRequestsQ
	FilterByMarketplaceId(id ...int64) NftRequestsQ

	FilterUpdateByMarketplaceId(id ...int64) NftRequestsQ
	UpdateStatus(newStatus resources.NftRequestStatus) NftRequestsQ
	Update() error
}

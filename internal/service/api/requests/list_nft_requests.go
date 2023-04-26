package requests

import (
	"github.com/pkg/errors"
	"net/http"

	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type ListNftRequestsRequest struct {
	pgdb.OffsetPageParams
	Sorts             pgdb.Sorts                   `url:"sort" default:"id"`
	PayerAddress      []string                     `filter:"payer_address"`
	CollectionAddress []string                     `filter:"collection_address"`
	NftId             []int64                      `filter:"nft_id"`
	ChainId           []int64                      `filter:"chain_id"`
	BookId            []int64                      `filter:"book_id"`
	Status            []resources.NftRequestStatus `filter:"status"`
}

func NewListNftRequestsRequest(r *http.Request) (*ListNftRequestsRequest, error) {
	var request ListNftRequestsRequest

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list promocodes request")
	}

	return &request, nil
}

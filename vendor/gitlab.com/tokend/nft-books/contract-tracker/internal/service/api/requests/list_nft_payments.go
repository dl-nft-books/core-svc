package requests

import (
	"github.com/pkg/errors"
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type ListNftPaymentsRequest struct {
	pgdb.OffsetPageParams
	Sorts pgdb.Sorts `url:"sort" default:"id"`

	Id         []int64  `filter:"id"`
	BookId     []int64  `filter:"book_id"`
	NftAddress []string `filter:"nft_address"`
	NftId      []int64  `filter:"nft_id"`
}

func NewListNftPaymentsRequest(r *http.Request) (*ListNftPaymentsRequest, error) {
	var request ListNftPaymentsRequest

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	return &request, nil
}

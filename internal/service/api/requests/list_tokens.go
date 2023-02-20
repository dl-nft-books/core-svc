package requests

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type ListTokensRequest struct {
	pgdb.OffsetPageParams
	Sorts pgdb.Sorts `url:"sort" default:"id"`
	
	Account        []string                `filter:"account"`
	Status         []resources.TokenStatus `filter:"status"`
	TokenId        *int64                  `filter:"token_id""`
	ChainId        []int64                 `filter:"chain_id""`
	MetadataHash   []string                `filter:"metadata_hash"`
	Name           []string                `filter:"name"`
	IsTokenPayment *bool                   `filter:"is_token_payment"`
}

func NewListTokensRequest(r *http.Request) (*ListTokensRequest, error) {
	var request ListTokensRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list promocodes request")
	}

	return &request, nil
}

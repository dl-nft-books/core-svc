package requests

import (
	"github.com/pkg/errors"
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type ListPromocodesRequest struct {
	pgdb.OffsetPageParams
	Sorts pgdb.Sorts `url:"sort" default:"id"`

	State []resources.PromocodeState `filter:"state"`
}

func NewListPromocodesRequest(r *http.Request) (*ListPromocodesRequest, error) {
	var request ListPromocodesRequest

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list promocodes request")
	}

	return &request, nil
}

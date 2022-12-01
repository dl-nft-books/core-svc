package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type ListPaymentsRequest struct {
	pgdb.OffsetPageParams
	Sorts pgdb.Sorts `url:"sort" default:"id"`

	Id           []int64  `filter:"id"`
	BookId       []int64  `filter:"book_id"`
	TokenAddress []string `filter:"token_address"`
}

func NewListPaymentsRequest(r *http.Request) (*ListPaymentsRequest, error) {
	var request ListPaymentsRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

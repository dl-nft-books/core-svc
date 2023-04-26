package requests

import (
	"github.com/pkg/errors"
	"net/http"

	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type ListTasksRequest struct {
	pgdb.OffsetPageParams
	Sorts pgdb.Sorts `url:"sort" default:"id"`

	Account  *string               `filter:"account"`
	Status   *resources.TaskStatus `filter:"status"`
	IpfsHash *string               `filter:"ipfs_hash"`
	TokenId  *int64                `filter:"token_id"`
}

func NewListTasksRequest(r *http.Request) (*ListTasksRequest, error) {
	var request ListTasksRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list task request")
	}

	return &request, nil
}

package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
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
		return nil, err
	}

	return &request, nil
}

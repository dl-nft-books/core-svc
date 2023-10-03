package responses

import (
	booker "github.com/dl-nft-books/book-svc/connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewNftRequestListResponse(r *http.Request, request *requests.ListNftRequestsRequest, nftRequests []data.NftRequest, booksApi booker.Connector) (*resources.NftRequestListResponse, error) {
	response := resources.NftRequestListResponse{}

	if len(nftRequests) == 0 {
		return &resources.NftRequestListResponse{
			Data: make([]resources.NftRequest, 0),
		}, nil
	}

	for _, nftRequest := range nftRequests {
		response.Data = append(response.Data, nftRequest.Resource())

		bookResponse, err := booksApi.GetBookById(nftRequest.BookId, nftRequest.ChainId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get book by its id", logan.F{
				"book_id": nftRequest.BookId,
			})
		}

		response.Included.Add(convertBookToResource(*bookResponse))
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}

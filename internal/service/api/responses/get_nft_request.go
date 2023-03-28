package responses

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewGetNftRequestResponse(nftRequest data.NftRequest) *resources.NftRequestResponse {
	return &resources.NftRequestResponse{Data: nftRequest.Resource()}
}

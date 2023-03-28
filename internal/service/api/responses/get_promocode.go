package responses

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewGetPromocodeResponse(promocode data.Promocode) *resources.PromocodeResponse {
	return &resources.PromocodeResponse{Data: promocode.Resource()}
}

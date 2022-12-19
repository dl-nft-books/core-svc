package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewGetPromocodeResponse(promocode data.Promocode) (*resources.PromocodeResponse, error) {
	return &resources.PromocodeResponse{Data: promocode.Resource()}, nil
}

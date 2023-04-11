package responses

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewGetPromocodeResponse(promocode data.Promocode) (*resources.PromocodeResponse, error) {
	response, err := promocode.Resource()
	if err != nil {
		return nil, err
	}
	return &resources.PromocodeResponse{Data: *response}, nil
}

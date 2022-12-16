package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewGetPromocodeResponse(promocode data.Promocode) (*resources.PromocodeResponse, error) {
	response := resources.PromocodeResponse{}

	response.Data = promocode.Resource()

	return &response, nil
}

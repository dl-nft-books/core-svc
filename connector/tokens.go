package connector

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/connector/models"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const tokensEndpoint = "tokens"

func (c *Connector) CreateToken(params models.CreateTokenParams) (id int64, err error) {
	var (
		response resources.KeyResponse
		request  = resources.CreateToken{
			Key: resources.NewKeyInt64(0, resources.TOKENS),
			Attributes: resources.CreateTokenAttributes{
				Account:      params.Account,
				MetadataHash: params.MetadataHash,
				Status:       params.Status,
				TokenId:      params.TokenId,
			},
		}
	)

	endpoint := fmt.Sprintf("%s/%s", c.baseUrl, tokensEndpoint)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return 0, errors.Wrap(err, "failed to marshal request")
	}

	if err = c.post(endpoint, requestAsBytes, &response); err != nil {
		return 0, errors.Wrap(err, "failed to create token")
	}

	createdTokenId := cast.ToInt64(response.Data.ID)
	return createdTokenId, nil
}

func (c *Connector) UpdateToken(params models.UpdateTokenParams) error {
	request := resources.UpdateToken{
		Key: resources.NewKeyInt64(params.Id, resources.TOKENS),
		Attributes: resources.UpdateTokenAttributes{
			Owner:   params.Owner,
			Status:  params.Status,
			TokenId: params.TokenId,
		},
	}

	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, tokensEndpoint, request.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListTokens(request models.ListTokensRequest) (*models.ListTokensResponse, error) {
	var result models.ListTokensResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s?%s", c.baseUrl, tokensEndpoint, urlval.MustEncode(request))

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetTokenById(id int64) (*models.TokenResponse, error) {
	var result models.TokenResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, tokensEndpoint, id)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}

	return &result, nil
}

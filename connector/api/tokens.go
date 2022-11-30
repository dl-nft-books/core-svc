package api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const tokensEndpoint = "tokens"

func (c *Connector) CreateToken(request resources.CreateToken) (id int64, err error) {
	var response resources.KeyResponse

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

func (c *Connector) UpdateToken(request resources.UpdateToken) error {
	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, tokensEndpoint, request.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListTokens(request requests.ListTokensRequest) (*resources.TokenListResponse, error) {
	var result resources.TokenListResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s?%s", c.baseUrl, tokensEndpoint, urlval.MustEncode(request))

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetTokenById(id int64) (*resources.TokenResponse, error) {
	var result resources.TokenResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, tokensEndpoint, id)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}

	return &result, nil
}

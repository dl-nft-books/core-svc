package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
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

package api

import (
	"errors"
	"fmt"

	"gitlab.com/tokend/nft-books/blob-svc/connector/models"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
)

func (c *Connector) GetReadableLink(key string) (models.LinkResponse, error) {
	var response models.LinkResponse

	exists, err := helpers.IsKeyExists(key, c.awsParams)
	if err != nil {
		return response, err
	}
	if !exists {
		return response, errors.New("key is not exist")
	}

	response.Data.Attributes.Url = fmt.Sprintf(
		"https://%s.s3.%s.amazonaws.com/%s",
		c.awsParams.Bucket,
		c.awsParams.Region,
		key,
	)

	return response, nil
}

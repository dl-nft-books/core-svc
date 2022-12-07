package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Connector struct {
	token   string
	baseUrl string
	client  *http.Client
}

func NewConnector(authToken, serviceUrl string) *Connector {
	return &Connector{
		token:   authToken,
		baseUrl: serviceUrl,
		client:  http.DefaultClient,
	}
}

func (c *Connector) get(endpoint string, dst interface{}) (found bool, err error) {
	// creating request
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, errors.Wrap(err, "failed to create connector request")
	}

	// setting headers
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// sending request
	response, err := c.client.Do(request)
	if err != nil {
		return false, errors.Wrap(err, "failed to process request")
	}

	if response.StatusCode == http.StatusNotFound {
		return false, nil
	}

	defer func(Body io.ReadCloser) {
		if tempErr := Body.Close(); tempErr != nil {
			err = tempErr
		}
	}(response.Body)

	// parsing response
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, errors.Wrap(err, "failed to read response body")
	}

	return json.Unmarshal(raw, &dst) == nil, err
}

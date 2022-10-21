package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type PriceResponse struct {
	Data struct {
		Attributes struct {
			Price string `json:"price"`
		} `json:"attributes"`
	} `json:"data"`
}

func GetPrice(r *http.Request, tokenAddress, platform string) (string, error) {
	coingeckoConfig := Coingecko(r)

	params := url.Values{}
	if tokenAddress != "" {
		params.Add("contract", tokenAddress)
	}
	params.Add("platform", platform)

	raw, err := sendRequest(params, coingeckoConfig.Host, coingeckoConfig.Endpoint)
	if err != nil {
		return "", err
	}

	return getPrice(raw)
}

func sendRequest(params url.Values, host, endpoint string) ([]byte, error) {
	rawQuery := params.Encode()

	url := fmt.Sprintf("http://%s/%s?%s",
		host,
		endpoint,
		rawQuery,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getPrice(raw []byte) (string, error) {
	var res PriceResponse

	if err := json.Unmarshal(raw, &res); err != nil {
		return "", err
	}

	return res.Data.Attributes.Price, nil
}

func ConvertPrice(raw string, precision int) (*big.Int, error) {
	floatPriceParts := strings.Split(raw, ".")
	if len(floatPriceParts) != 2 {
		return nil, errors.New("invalid price response")
	}

	// just gluing two parts
	stringPrice := floatPriceParts[0] + floatPriceParts[1]
	// minus the len of the float part
	precision -= len(floatPriceParts[1])

	price, ok := big.NewInt(0).SetString(fmt.Sprintf("%v", stringPrice), 10)
	if !ok {
		return nil, errors.New("failed to set bigint value")
	}

	price.Mul(price, big.NewInt(int64(math.Pow10(precision))))

	return price, nil
}

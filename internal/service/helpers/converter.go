package helpers

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

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

	if precision < 0 {
		positivePrecision := int(math.Abs(float64(precision)))
		price.Div(price, big.NewInt(int64(math.Pow10(positivePrecision))))
	} else {
		price.Mul(price, big.NewInt(int64(math.Pow10(precision))))
	}

	return price, nil
}

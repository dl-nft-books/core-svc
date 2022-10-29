package helpers

import (
	"strings"
	"testing"
)

const (
	defaultPrecision         = 18
	defaultResponsePrecision = 6
	msgInvalidResult         = "Invalid result"
)

func TestConverter(t *testing.T) {
	var price, expected, actual string

	t.Run("Default conversion 1", func(t *testing.T) {
		price = "1299.00000"

		logInput(t, price, defaultPrecision)

		res, err := ConvertPrice(price, defaultPrecision)
		failOnError(t, err)

		expected = "1299" + strings.Repeat("0", defaultPrecision)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Default conversion 2", func(t *testing.T) {
		price = "178411.856734"

		logInput(t, price, defaultPrecision)

		res, err := ConvertPrice(price, defaultPrecision)
		failOnError(t, err)

		expected = "178411856734" + strings.Repeat("0", defaultPrecision-defaultResponsePrecision)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Default conversion 3", func(t *testing.T) {
		price = "0001.000001"

		logInput(t, price, defaultPrecision)

		res, err := ConvertPrice(price, defaultPrecision)
		failOnError(t, err)

		expected = "1000001" + strings.Repeat("0", defaultPrecision-defaultResponsePrecision)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Default conversion 4", func(t *testing.T) {
		price = "0.555555"

		logInput(t, price, defaultPrecision)

		res, err := ConvertPrice(price, defaultPrecision)
		failOnError(t, err)

		expected = "555555" + strings.Repeat("0", defaultPrecision-defaultResponsePrecision)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Change Precision 1", func(t *testing.T) {
		price = "123.000001"

		newPrecision := defaultPrecision - 3

		logInput(t, price, newPrecision)

		res, err := ConvertPrice(price, newPrecision)
		failOnError(t, err)

		expected = "123000001" + strings.Repeat("0", newPrecision-defaultResponsePrecision)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Change Precision 2", func(t *testing.T) {
		price = "1.000001"

		newPrecision := defaultPrecision - 10

		logInput(t, price, newPrecision)

		res, err := ConvertPrice(price, newPrecision)
		failOnError(t, err)

		expected = "1000001" + strings.Repeat("0", newPrecision-defaultResponsePrecision)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Change Precision 3", func(t *testing.T) {
		price = "1.124395"

		newPrecision := 4

		logInput(t, price, newPrecision)

		res, err := ConvertPrice(price, newPrecision)
		failOnError(t, err)

		expected = "11243"
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Change response precision 1", func(t *testing.T) {
		price = "1.23"

		logInput(t, price, defaultPrecision)

		res, err := ConvertPrice(price, defaultPrecision)
		failOnError(t, err)

		expected = "123" + strings.Repeat("0", defaultPrecision-2)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Change response precision 2", func(t *testing.T) {
		price = "1.2345"

		logInput(t, price, defaultPrecision)

		res, err := ConvertPrice(price, defaultPrecision)
		failOnError(t, err)

		expected = "12345" + strings.Repeat("0", defaultPrecision-4)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})

	t.Run("Change default and response precision", func(t *testing.T) {
		price = "1.234"
		newPrec := defaultPrecision - 4

		logInput(t, price, newPrec)

		res, err := ConvertPrice(price, newPrec)
		failOnError(t, err)

		expected = "1234" + strings.Repeat("0", newPrec-3)
		actual = res.String()

		logOutput(t, expected, actual)

		assert(t, expected, actual, msgInvalidResult)
	})
}

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assert(t *testing.T, expected, actual interface{}, msg string) {
	if expected != actual {
		if actual == nil {
			actual = "<nil> (not found or invalid type)"
		}
		t.Fatalf("%s: expected %v, actual %v", msg, expected, actual)
	}
}

func logInput(t *testing.T, price, precision interface{}) {
	t.Logf("\nPrice: %s\nPrecision: %v", price, precision)
}

func logOutput(t *testing.T, exp, act interface{}) {
	t.Logf("\nExpected:	%s\nActual:		%v", exp, act)
}

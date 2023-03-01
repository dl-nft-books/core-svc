package helpers

import (
	"math"
)

func Trancate(n float64, decimal int) float64 {
	return math.Floor(n*math.Pow10(decimal)) / math.Pow10(decimal)
}

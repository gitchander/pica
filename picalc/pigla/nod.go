package pigla

import (
	"math"
)

// NumberOfDigits
// RepresentLen

// NumberOfDigits returns number of digits needed to represent value in base
func NumberOfDigits(base, value float64) float64 {
	return math.Log(value)/math.Log(base) + 1
}

func invNumberOfDigits(base, nod float64) (value float64) {
	return math.Pow(base, nod-1)
}

func ConvertNOD(base1, nod1 float64, base2 float64) (nod2 float64) {
	t := math.Log(base1) / math.Log(base2)
	nod2 = (nod1-1)*t + 1
	return nod2
}

package pigla

import (
	"math"
)

// Gauss–Legendre algorithm
// https://en.wikipedia.org/wiki/Gauss%E2%80%93Legendre_algorithm

func CalcPi(n int) float64 {

	var (
		a = 1.0
		b = math.Sqrt2 / 2
		t = 0.25 // 1.0 / 4.0
		p = 1.0
	)

	for i := 0; i < n; i++ {

		a_next := (a + b) / 2
		b_next := math.Sqrt(a * b)
		t_next := t - p*square(a-a_next)
		p_next := 2 * p

		a = a_next
		b = b_next
		t = t_next
		p = p_next
	}

	Pi := square((a+b)/2) / t

	return Pi
}

func square(a float64) float64 {
	return a * a
}

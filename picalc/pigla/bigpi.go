package pigla

import (
	"math/big"
)

// Gauss–Legendre algorithm
// https://en.wikipedia.org/wiki/Gauss%E2%80%93Legendre_algorithm

type StepFunc func(step int, pi *big.Float)

func CalcBigPi(prec uint, steps int, sf StepFunc) *big.Float {

	var (
		one  = newBigFloatPrec(prec).SetInt64(1)
		two  = newBigFloatPrec(prec).SetInt64(2)
		four = newBigFloatPrec(prec).SetInt64(4)

		sqrt2 = newBigFloatPrec(prec).Sqrt(two) // sqrt(2)
	)

	var (
		a = newBigFloatPrec(prec).SetInt64(1)     // a = 1
		b = newBigFloatPrec(prec).Quo(one, sqrt2) // b = 1 / sqrt(2)
		t = newBigFloatPrec(prec).Quo(one, four)  // t = 1.0 / 4.0
		p = newBigFloatPrec(prec).SetInt64(1)     // p = 1
	)

	var (
		an = new(big.Float)
		bn = new(big.Float)

		v  = new(big.Float)
		t4 = new(big.Float) // t * 4

		pi = new(big.Float)
	)

	for i := 0; i < steps; i++ {

		an.Copy(a) // an = a
		bn.Copy(b) // bn = b

		// a = (an + bn) / 2
		v.Add(an, bn)
		a.Quo(v, two)

		// b = sqrt(an * bn)
		v.Mul(an, bn)
		b.Sqrt(v)

		// t = t - p * sqr(an - a)
		v.Sub(an, a)
		v.Mul(v, v)
		v.Mul(v, p)
		t.Sub(t, v)

		// p = 2 * p
		p.Mul(two, p)

		if sf != nil {
			// Pi = sqr(a + b) / (4 * t)
			v.Add(a, b)
			v.Mul(v, v)
			t4.Mul(four, t)
			pi.Quo(v, t4)

			sf(i, pi)
		}
	}

	// Pi = sqr(a + b) / (4 * t)
	v.Add(a, b)
	v.Mul(v, v)
	t4.Mul(four, t)
	pi.Quo(v, t4)

	return pi
}

func newBigFloatPrec(prec uint) *big.Float {
	return new(big.Float).SetPrec(prec)
}

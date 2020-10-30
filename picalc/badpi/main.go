package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	r := Rect{
		Min: Pt2f(0, 0),
		Max: Pt2f(1, 1),
	}

	lo, hi := iter(r, 20)

	lo *= 4
	hi *= 4

	fmt.Printf("%.15f\n", hi-lo)
	fmt.Println((lo + hi) / 2)

	// real Pi = 3.141592653589793238462643383279...

	fmt.Println(time.Since(start))
}

func iter(r Rect, n int) (lo float64, hi float64) {

	dMin := r.Min.Norm()
	if dMin > 1 {
		return 0, 0
	}

	dMax := r.Max.Norm()
	if dMax < 1 {
		a := r.Area()
		return a, a
	}

	if n == 0 {
		return 0, r.Area()
	}

	rs := r.Split()
	for _, r := range rs {
		lo1, hi1 := iter(r, n-1)

		lo += lo1
		hi += hi1
	}

	return lo, hi
}

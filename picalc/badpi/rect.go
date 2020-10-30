package main

type Rect struct {
	Min Point2f
	Max Point2f
}

func (r Rect) Area() float64 {
	p := r.Max.Sub(r.Min)
	return p.X * p.Y
}

func (r Rect) Split() []Rect {
	middle := r.Min.Add(r.Max).DivScalar(2)
	return []Rect{
		{
			Min: r.Min,
			Max: middle,
		},
		{
			Min: middle,
			Max: r.Max,
		},
		{
			Min: Pt2f(r.Min.X, middle.Y),
			Max: Pt2f(middle.X, r.Max.Y),
		},
		{
			Min: Pt2f(middle.X, r.Min.Y),
			Max: Pt2f(r.Max.X, middle.Y),
		},
	}
}

package geometricForm

import "math"

type GeometricForm interface {
	calculateArea() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Square struct {
	Side float64
}

func (s Square) CalculateArea() float64 {
	return math.Pow(s.Side, 2)
}

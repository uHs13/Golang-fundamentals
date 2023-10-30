package main

import (
	"fmt"
	"math"
)

func main() {
	circle := circle{radius: 5}
	square := square{side: 10}

	fmt.Println(calculateArea(circle))
	fmt.Println(calculateArea(square))
}

func calculateArea(form geometricForm) float64 {
	return form.calculateArea()
}

type geometricForm interface {
	calculateArea() float64
}

type circle struct {
	radius float64
}

func (c circle) calculateArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type square struct {
	side float64
}

func (s square) calculateArea() float64 {
	return math.Pow(s.side, 2)
}

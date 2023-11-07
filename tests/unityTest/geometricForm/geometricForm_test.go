package geometricForm_test

import (
	"math"
	"testing"
	. "unity-tests/geometricForm"
)

func TestShouldCorrectlyCalculateArea(t *testing.T) {
	t.Run("Circle", func(t *testing.T) {
		radius := float64(10)
		circle := Circle{radius}

		expectedArea := math.Pi * math.Pow(radius, 2)
		area := circle.CalculateArea()

		if area != expectedArea {
			t.Fatalf(
				"Area %f is not equal to expected %f",
				area,
				expectedArea,
			)
		}
	})

	t.Run("Square", func(t *testing.T) {
		side := float64(10)
		square := Square{side}

		expectedArea := math.Pow(10, 2)
		area := square.CalculateArea()

		if area != expectedArea {
			t.Fatalf(
				"Area %f is not equal to expected %f",
				area,
				expectedArea,
			)
		}
	})
}

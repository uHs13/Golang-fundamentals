package main

import "fmt"

func main() {
	var positiveNumber float64 = 100

	invertedNumber := changeNumberSignalWithoutPointer(float64(positiveNumber))

	fmt.Println(invertedNumber)
	fmt.Println(positiveNumber)

	changeNumberSignalWithPointer(&positiveNumber)

	fmt.Println(positiveNumber)
}

func changeNumberSignalWithoutPointer(number float64) float64 {
	return number * -1
}

func changeNumberSignalWithPointer(number *float64) {
	*number = *number * -1
}

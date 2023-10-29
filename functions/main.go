package main

import "fmt"

func main() {
	fmt.Println(addNumbers(10, 15))

	var function = func(integerOne int, integerTwo int) int {
		return integerOne * integerTwo
	}

	fmt.Println(function(10, 5))

	addedNumbers, subtractedNumbers, multipliedNumbers, dividedNumbers := mathematicalCalculationsWithMoreThanOneReturnedValue(50, 10)

	fmt.Println(addedNumbers)
	fmt.Println(subtractedNumbers)
	fmt.Println(multipliedNumbers)
	fmt.Println(dividedNumbers)

	returnOne, _ := ignoreReturn()

	fmt.Println(returnOne)
}

func addNumbers(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

func mathematicalCalculationsWithMoreThanOneReturnedValue(numberOne int, numberTwo int) (int, int, float64, float64) {
	addNumbers := numberOne + numberTwo
	subtractNumbers := numberOne - numberTwo
	multiplyNumbers := numberOne * numberTwo
	divideNumbers := numberOne / numberTwo

	return addNumbers, subtractNumbers, float64(multiplyNumbers), float64(divideNumbers)
}

func ignoreReturn() (string, string) {
	return "a", "b"
}

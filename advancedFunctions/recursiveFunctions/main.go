package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	factorial := number * factorial(number-1)

	return factorial
}

func main() {
	factorial := factorial(3)

	fmt.Println(factorial)
}

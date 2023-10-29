package main

import "fmt"

func main() {
	fmt.Println(panicProgramIfAverageNumberIsLowerThanTen(16, 3))
}

func panicProgramIfAverageNumberIsLowerThanTen(numberOne, numberTwo float64) string {
	defer recoverProgram()

	average := (numberOne + numberTwo) / 2

	if average < 10 {
		panic("The average number is lower than ten")
	}

	return fmt.Sprintf("The average number %f is greater than 10", average)
}

func recoverProgram() {
	if recover := recover(); recover != nil {
		fmt.Println("Program execution successfully recovered")
	}
}

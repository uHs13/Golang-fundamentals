package main

import "fmt"

func calculations(numberOne int, numberTwo int) (sum int, sub int) {
	sum = numberOne + numberTwo
	sub = numberOne - numberTwo

	return
}

func main() {
	sum, sub := calculations(10, 5)

	fmt.Println(sum, sub)
}

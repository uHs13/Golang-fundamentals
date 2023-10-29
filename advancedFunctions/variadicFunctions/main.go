package main

import "fmt"

func sum(numbers ...int) int {
	var sum int

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func main() {
	sum := sum(1, 2, 3, 4, 5)

	fmt.Println(sum)
}

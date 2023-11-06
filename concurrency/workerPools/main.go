package main

import "fmt"

func main() {
	length := 42
	numbers := make(chan int, length)
	results := make(chan int, length)

	go worker(numbers, results)
	go worker(numbers, results)
	go worker(numbers, results)

	for counter := 0; counter < length; counter++ {
		numbers <- counter
	}

	close(numbers)

	for counter := 0; counter < length; counter++ {
		result := <-results

		fmt.Println(result)
	}
}

func worker(numbers <-chan int, results chan<- int) {
	for number := range numbers {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

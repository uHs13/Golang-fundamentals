package main

import (
	"fmt"
	"math/rand"
)

const limit int = 100

func main() {
	var number int = rand.Intn(200)
	var result string

	if isNumberGreaterThanLimit(number) {
		result = fmt.Sprintf("%d is greater than %d", number, limit)
	} else {
		result = fmt.Sprintf("%d is not greater than %d", number, limit)
	}

	fmt.Println(result)

	// if initialization statement

	if randomNumber := rand.Intn(500); randomNumber > limit {
		result = fmt.Sprintf("%d is greater than %d", randomNumber, limit)
	}

	fmt.Println(result)
}

func isNumberGreaterThanLimit(number int) bool {
	return number > limit
}

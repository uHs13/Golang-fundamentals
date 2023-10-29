package main

import (
	"fmt"
)

func main() {
	var counter int = 0

	for counter <= 20 {
		fmt.Println(counter)
		counter++
	}

	for counter = 0; counter < 10; counter++ {
		fmt.Println(counter)
	}

	names := [3]string{"Jhon", "Carl", "Doc"}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, letter := range "golang" {
		fmt.Println(string(letter))
	}

	user := map[string]string{
		"name":  "John Cena",
		"email": "john.cena@email.com",
	}

	for index, value := range user {
		fmt.Println(index, value)
	}
}

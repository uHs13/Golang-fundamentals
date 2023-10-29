package main

import "fmt"

func main() {
	closure := closure()

	closure()
}

func closure() func() {
	var text string = "The orange is orange"

	return func() {
		fmt.Println(text)
	}
}

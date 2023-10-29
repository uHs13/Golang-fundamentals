package main

import "fmt"

func main() {
	anonymous := func(text string, number float64) string {
		return fmt.Sprintf("%s -> %f", text, number)
	}("Some text", 13.3)

	fmt.Println(anonymous)
}

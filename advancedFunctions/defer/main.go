package main

import "fmt"

func printSomeText() {
	fmt.Println("Some text")
}

func printSomePhrase() {
	fmt.Println("The grass is green")
}

func main() {
	defer printSomeText()

	printSomePhrase()
}

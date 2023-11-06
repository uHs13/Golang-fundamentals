package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Some words"
	channel <- "Another words"

	message := <-channel
	messageTwo := <-channel

	fmt.Println(message)
	fmt.Println(messageTwo)
}

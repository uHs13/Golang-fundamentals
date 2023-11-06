package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("The air is fresh today", channel)

	for message := range channel {
		fmt.Println(message)
	}
}

func write(text string, channel chan string) {
	for counter := 0; counter < 5; counter++ {
		channel <- text

		time.Sleep(time.Second)
	}

	close(channel)
}

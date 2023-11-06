package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplexing(
		write("John Cena"),
		write("Carl Johnson"),
	)

	for counter := 0; counter < 10; counter++ {
		fmt.Println(<-channel)
	}
}

func multiplexing(
	channelOne <-chan string,
	channelTwo <-chan string,
) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelOne:
				outputChannel <- message
			case message := <-channelTwo:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("received value: %s", text)

			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

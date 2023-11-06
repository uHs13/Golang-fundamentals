package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("John Cena")

	for counter := 0; counter < 15; counter++ {
		fmt.Println(<-channel)
	}
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

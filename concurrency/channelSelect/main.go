package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)

			channelOne <- "message 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)

			channelTwo <- "message 2"
		}
	}()

	for {
		select {
		case messageOne := <-channelOne:
			fmt.Println(messageOne)
		case messageTwo := <-channelTwo:
			fmt.Println(messageTwo)
		}
	}
}

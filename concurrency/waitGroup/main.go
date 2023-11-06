package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("I have seen a big dog")

		waitGroup.Done()
	}()

	go func() {
		write("I have seen a big cat")

		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for counter := 0; counter < 5; counter++ {
		fmt.Println(text)

		time.Sleep(time.Second)
	}
}

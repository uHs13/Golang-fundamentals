package main

import (
	"fmt"
	"time"
)

func main() {
	go write("John Cena")
	write("Green")
}

func write(text string) {
	for {
		fmt.Println(text)

		time.Sleep(time.Second)
	}
}

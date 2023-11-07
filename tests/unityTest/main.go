package main

import (
	"fmt"
	"unity-tests/address"
)

func main() {
	completeAddress := "Avenida Madrid"

	isValidAddress := address.IsValidBrazilianFormat(completeAddress)

	fmt.Println(isValidAddress)
}

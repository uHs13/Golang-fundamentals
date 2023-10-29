package main

import (
	"fmt"
	"module/secondaryPackage"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("First we gonna rock, then we gonna roll, then we let it pop, go let it go")
	secondaryPackage.WriteItDown()

	error := checkmail.ValidateFormat("123")

	fmt.Println(error)
}

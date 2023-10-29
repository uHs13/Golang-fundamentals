package main

import (
	"errors"
	"fmt"
)

func main() {
	var text string = "some text"

	anotherWayToDeclareString := "some text too"

	fmt.Println(text)
	fmt.Println(anotherWayToDeclareString)

	fmt.Println("// --- int --- //")

	var integerByComputerArchitecture int = 13
	var integerEight int8 = 99
	var integerSixteen int16 = 9999
	var integerThirtyTwo int32 = 999999999
	var integerSixtyFour int64 = 999999999999999999
	integerByInference := 123
	var negativeInteger int = -13

	var runeNumber rune = 98765 // same as var name int32

	fmt.Println(integerByComputerArchitecture)
	fmt.Println(integerEight)
	fmt.Println(integerSixteen)
	fmt.Println(integerThirtyTwo)
	fmt.Println(integerSixtyFour)
	fmt.Println(integerByInference)
	fmt.Println(negativeInteger)
	fmt.Println(runeNumber)

	fmt.Println("// --- uint --- //")

	var unsignedIntegerByComputerArchitecture uint = 13
	var uintEight uint8 = 99
	var uintSixteen uint16 = 9999
	var uintThirtyTwo uint32 = 999999999
	var uintSixtyFour uint64 = 999999999999999999
	uintByInference := 123

	var byteNumber byte = 123 // same as var name uint8

	fmt.Println(unsignedIntegerByComputerArchitecture)
	fmt.Println(uintEight)
	fmt.Println(uintSixteen)
	fmt.Println(uintThirtyTwo)
	fmt.Println(uintSixtyFour)
	fmt.Println(uintByInference)
	fmt.Println(byteNumber)

	fmt.Println("// --- float --- //")

	var realNumber32 float32 = 123.45
	var realNumber64 float64 = 678.910
	realNumberByInference := 456.78

	fmt.Println(realNumber32)
	fmt.Println(realNumber64)
	fmt.Println(realNumberByInference)

	fmt.Println("// --- string --- //")

	var some string = "some text"
	stringByInference := "some text two"

	fmt.Println(some)
	fmt.Println(stringByInference)

	asciiTableValue := 'C'

	fmt.Println(asciiTableValue)

	fmt.Println("// --- bool --- //")

	var trueVariable bool = true
	var falseVariable bool = false
	booleanByInference := false

	fmt.Println(trueVariable)
	fmt.Println(falseVariable)
	fmt.Println(booleanByInference)

	fmt.Println("// --- error --- //")

	var errorVariable error
	var errorTwo error = errors.New("Internal server error")

	fmt.Println(errorVariable)
	fmt.Println(errorTwo)
}

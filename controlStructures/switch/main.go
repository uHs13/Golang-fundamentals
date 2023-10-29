package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var result string = generatePhrase(rand.Intn(4))

	fmt.Println(result)
}

func generatePhrase(weekDayNumber int) string {
	availableWeekDays := [4]string{"Monday", "Tuesday", "Wednesday", "Thursday"}

	weekDayName := availableWeekDays[weekDayNumber]

	switch weekDayName {
	case "Monday":
		return "Monday left me broken"
	case "Tuesday":
		return "Tuesday I was through with hoping"
	case "Wednesday":
		return "Wednesday my empty arms were open"
	case "Thursday":
		return "Thursday waiting for love"
	default:
		return "Unfortunately nothing happened"
	}
}

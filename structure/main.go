package main

import "fmt"

type address struct {
	address string
	zipCode int
}

type user struct {
	name    string
	age     int
	email   string
	address address
}

type employee struct {
	user
	serialNumber string
	sector       string
}

func main() {
	var userOne user

	userOne.name = "John Cena"
	userOne.age = 90
	userOne.email = "john.cena@email.com"

	fmt.Println(userOne)

	address := address{"some street", 1234}

	userTwo := user{"John Wick", 50, "john.wick@email.com", address}

	fmt.Println(userTwo)

	userThree := user{name: "John Titor", email: "john.titor@email.com"}

	fmt.Println(userThree)

	fmt.Println("Something in Go that is like heritage/composition")

	employee := employee{user: userThree, serialNumber: "9e87iw", sector: "IT"}

	fmt.Println(employee.name)

	fmt.Println("Pointers")

	var number int = 1000
	var pointer *int = &number

	number++

	fmt.Println(*pointer)
}

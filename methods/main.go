package main

import "fmt"

func main() {
	user := user{name: "John Cena", age: 35, email: "john.cena@email.com"}

	fmt.Println(user.getName())
	fmt.Println(user.getAge())
	fmt.Println(user.getEmail())
}

type user struct {
	name  string
	age   uint
	email string
}

func (u user) getName() string {
	return u.name
}

func (u user) getAge() uint {
	return u.age
}

func (u user) getEmail() string {
	return u.email
}

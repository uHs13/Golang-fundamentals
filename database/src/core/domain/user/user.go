package userDomain

type User struct {
	Id    string
	Name  string
	Email string
}

func NewUser(id string, name string, email string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

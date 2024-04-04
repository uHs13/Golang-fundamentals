package userDomain

type User struct {
	Id          string
	Name        string
	Email       string
	DateCreated string
	DateUpdated string
	DateDeleted string
}

func NewUser(
	id string,
	name string,
	email string,
	dateCreated string,
) *User {
	return &User{
		Id:          id,
		Name:        name,
		Email:       email,
		DateCreated: dateCreated,
	}
}

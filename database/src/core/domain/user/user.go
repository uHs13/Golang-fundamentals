package userDomain

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateCreated string `json:"-"`
	DateUpdated string `json:"-"`
	DateDeleted string `json:"-"`
}

func NewUser(
	id string,
	name string,
	email string,
	dateCreated string,
	dateUpdated string,
	dateDeleted string,
) *User {
	return &User{
		Id:          id,
		Name:        name,
		Email:       email,
		DateCreated: dateCreated,
		DateUpdated: dateUpdated,
		DateDeleted: dateDeleted,
	}
}

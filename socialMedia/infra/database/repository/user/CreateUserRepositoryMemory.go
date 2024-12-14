package userrepository

import (
	"socialMedia/application/domain/user"
	port "socialMedia/application/port/repository/user"
)

type CreateUserRepositoryMemory struct {
}

func NewCreateUserRepositoryMemory() port.CreateUserRepositoryInterface {
	return &CreateUserRepositoryMemory{}
}

func (repository *CreateUserRepositoryMemory) Execute(user *user.User) (*user.User, error) {
	return user, nil
}

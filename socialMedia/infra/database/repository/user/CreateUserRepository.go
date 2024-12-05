package userRepository

import (
	"socialMedia/application/domain/user"
	port "socialMedia/application/port/database"
)

type CreateUserRepository struct {
	database port.DatabaseConnectionInterface
}

func NewCreateUserRepository(
	database port.DatabaseConnectionInterface,
) *CreateUserRepository {
	return &CreateUserRepository{
		database: database,
	}
}

func (repository *CreateUserRepository) Execute(user *user.User) (*user.User, error) {
	return nil, nil
}

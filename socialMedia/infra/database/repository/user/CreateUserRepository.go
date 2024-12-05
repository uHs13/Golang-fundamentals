package userRepository

import (
	"socialMedia/application/domain/user"
	port "socialMedia/application/port/database"
)

const (
	createUserQuery = "INSERT INTO user (id, name, nickname, email) VALUES (?, ?, ?, ?)"
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
	statement, err := repository.database.GetConnection().Prepare(createUserQuery)

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(
		user.GetId(),
		user.GetName(),
		user.GetNickname(),
		user.GetEmail(),
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

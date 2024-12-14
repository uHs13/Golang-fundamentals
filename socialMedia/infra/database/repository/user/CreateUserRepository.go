package userrepository

import (
	"socialMedia/application/domain/user"
	port "socialMedia/application/port/database"
	portRepository "socialMedia/application/port/repository/user"
)

const (
	createUserQuery = "INSERT INTO user (id, name, nickname, email, password) VALUES (?, ?, ?, ?, ?)"
)

type CreateUserRepository struct {
	database port.DatabaseConnectionInterface
}

func NewCreateUserRepository(
	database port.DatabaseConnectionInterface,
) portRepository.CreateUserRepositoryInterface {
	return &CreateUserRepository{
		database: database,
	}
}

func (repository *CreateUserRepository) Execute(user *user.User) (*user.User, error) {
	statement, err := repository.database.GetConnection().Prepare(createUserQuery)

	if err != nil {
		return nil, err
	}

	password := user.GetPassword()
	passwordHash, err := password.GenerateHash()

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(
		user.GetId(),
		user.GetName(),
		user.GetNickname(),
		user.GetEmail(),
		passwordHash,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

package repository

import (
	"database/sql"
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
)

type UserDatabase struct {
	connection *sql.DB
}

const (
	createUserQuery = "INSERT INTO user (id, name, email, created_at) VALUES (?, ?, ?, ?)"
)

func NewUserDatabase(connection *sql.DB) repositories.UserRepository {
	return &UserDatabase{
		connection: connection,
	}
}

func (userDatabase *UserDatabase) Create(user userDomain.User) error {
	statement, err := userDatabase.connection.Prepare(createUserQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(
		user.Id,
		user.Name,
		user.Email,
		user.DateCreated,
	)

	if err != nil {
		return err
	}

	return nil
}

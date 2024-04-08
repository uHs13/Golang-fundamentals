package repository

import (
	"database/sql"
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
	"fmt"
)

type UserDatabase struct {
	connection *sql.DB
}

const (
	createUserQuery                   = "INSERT INTO user (id_user, name, email, created_at) VALUES (?, ?, ?, ?)"
	isEmailAlreadyRegisteredUserQuery = "SELECT email FROM user WHERE email = ? LIMIT 1"
	isIdAlreadyRegisteredUserQuery    = "SELECT id_user FROM user WHERE id_user = ? LIMIT 1"
	findUsersQuery                    = "SELECT id_user, name, email FROM user WHERE deleted_at IS NULL ORDER BY created_at DESC"
	findUserQuery                     = "SELECT id_user, name, email FROM user WHERE id_user = ? AND deleted_at IS NULL"

	isEmailAlreadyRisteredErrorMsg = "invalid email"
	isIdAlreadyRisteredErrorMsg    = "invalid id"
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

func (userDatabase *UserDatabase) IsEmailAlreadyRegistered(user userDomain.User) error {
	statement, err := userDatabase.connection.Prepare(isEmailAlreadyRegisteredUserQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	rows, err := statement.Query(user.Email)

	if err != nil {
		return err
	}

	var userWithSameEmail userDomain.User

	for rows.Next() {
		if err := rows.Scan(&userWithSameEmail.Email); err != nil {
			return err
		}
	}

	if userWithSameEmail.Email != "" {
		return fmt.Errorf(isEmailAlreadyRisteredErrorMsg)
	}

	return nil
}

func (userDatabase *UserDatabase) IsIdAlreadyRegistered(user userDomain.User) error {
	statement, err := userDatabase.connection.Prepare(isIdAlreadyRegisteredUserQuery)

	if err != nil {
		return err
	}

	defer statement.Close()

	rows, err := statement.Query(user.Id)

	if err != nil {
		return err
	}

	var userWithSameId userDomain.User

	for rows.Next() {
		if err := rows.Scan(&userWithSameId.Id); err != nil {
			return err
		}
	}

	if userWithSameId.Id != "" {
		return fmt.Errorf(isIdAlreadyRisteredErrorMsg)
	}

	return nil
}

func (userDatabase *UserDatabase) FindUsers() ([]userDomain.User, error) {
	statement, err := userDatabase.connection.Prepare(findUsersQuery)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		return nil, err
	}

	var users []userDomain.User
	var user userDomain.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userDatabase *UserDatabase) FindUser(user userDomain.User) (userDomain.User, error) {
	statement, err := userDatabase.connection.Prepare(findUserQuery)

	if err != nil {
		return userDomain.User{}, err
	}

	defer statement.Close()

	rows, err := statement.Query(user.Id)

	if err != nil {
		return userDomain.User{}, err
	}

	var foundUser userDomain.User

	for rows.Next() {
		if err := rows.Scan(&foundUser.Id, &foundUser.Name, &foundUser.Email); err != nil {
			return userDomain.User{}, err
		}
	}

	return foundUser, nil
}

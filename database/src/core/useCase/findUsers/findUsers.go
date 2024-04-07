package findUsers

import (
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
)

type FindUsers struct {
	userRepository repositories.UserRepository
}

func NewFindUsers(userRepository repositories.UserRepository) *FindUsers {
	return &FindUsers{
		userRepository: userRepository,
	}
}

func (getUsers *FindUsers) Execute() ([]userDomain.User, error) {
	users, err := getUsers.
		userRepository.
		FindUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

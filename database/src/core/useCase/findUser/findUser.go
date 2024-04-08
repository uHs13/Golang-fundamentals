package findUser

import (
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
)

type FindUser struct {
	userRepository repositories.UserRepository
	user           userDomain.User
}

func NewFindUser(
	userRepository repositories.UserRepository,
	user userDomain.User,
) *FindUser {
	return &FindUser{
		userRepository: userRepository,
		user:           user,
	}
}

func (findUser *FindUser) Execute() (userDomain.User, error) {
	user, err := findUser.userRepository.FindUser(findUser.user)

	if err != nil {
		return userDomain.User{}, err
	}

	return user, nil
}

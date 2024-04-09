package findUser

import (
	"database/src/adapter/http/routesConstants"
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
	"fmt"
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

	if user == (userDomain.User{}) {
		return userDomain.User{}, fmt.Errorf(routesConstants.MessageUserNotFound)
	}

	return user, nil
}

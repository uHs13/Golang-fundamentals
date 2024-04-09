package editUser

import (
	"database/src/adapter/http/routesConstants"
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
	"fmt"
)

type EditUser struct {
	userRepository repositories.UserRepository
	user           userDomain.User
}

func NewEditUser(
	userRepository repositories.UserRepository,
	user userDomain.User,
) *EditUser {
	return &EditUser{
		userRepository: userRepository,
		user:           user,
	}
}

func (editUser *EditUser) Execute() error {
	user, err := editUser.userRepository.FindUser(editUser.user)

	if err != nil {
		return err
	}

	if user == (userDomain.User{}) {
		return fmt.Errorf(routesConstants.MessageUserNotFound)
	}

	err = editUser.
		userRepository.
		IsEmailAlreadyRegistered(editUser.user)

	if err != nil {
		return err
	}

	err = editUser.userRepository.Edit(editUser.user)

	if err != nil {
		return nil
	}

	return nil
}

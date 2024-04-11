package deleteUser

import (
	"database/src/adapter/http/routesConstants"
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
	"fmt"
)

type DeleteUser struct {
	userRepository repositories.UserRepository
	user           userDomain.User
}

func NewDeleteUser(
	userRepository repositories.UserRepository,
	user userDomain.User,
) *DeleteUser {
	return &DeleteUser{
		userRepository: userRepository,
		user:           user,
	}
}

func (deleteUser *DeleteUser) Execute() error {
	user, err := deleteUser.userRepository.FindUser(deleteUser.user)

	if err != nil {
		return err
	}

	if user == (userDomain.User{}) {
		return fmt.Errorf(routesConstants.MessageUserNotFound)
	}

	err = deleteUser.userRepository.Delete(deleteUser.user)

	if err != nil {
		return err
	}

	return nil
}

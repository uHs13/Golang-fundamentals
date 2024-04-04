package userUseCase

import (
	userDomain "database/src/core/domain/user"
	"database/src/core/port/repositories"
)

type CreateUser struct {
	userRepository repositories.UserRepository
	user           userDomain.User
}

func NewCreateUser(
	userRepository repositories.UserRepository,
	user userDomain.User,
) *CreateUser {
	return &CreateUser{
		userRepository: userRepository,
		user:           user,
	}
}

func (createUser *CreateUser) Execute() error {
	err := createUser.userRepository.Create(createUser.user)

	if err != nil {
		return err
	}

	return nil
}

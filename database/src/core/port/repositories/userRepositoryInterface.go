package repositories

import userDomain "database/src/core/domain/user"

type UserRepository interface {
	Create(user userDomain.User) error
	IsEmailAlreadyRegistered(user userDomain.User) error
	IsIdAlreadyRegistered(user userDomain.User) error
	FindUsers() ([]userDomain.User, error)
	FindUser(user userDomain.User) (userDomain.User, error)
	Edit(user userDomain.User) error
}

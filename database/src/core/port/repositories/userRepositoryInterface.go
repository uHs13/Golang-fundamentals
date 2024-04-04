package repositories

import userDomain "database/src/core/domain/user"

type UserRepository interface {
	Create(user userDomain.User) error
}

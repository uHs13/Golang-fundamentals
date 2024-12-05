package port

import "socialMedia/application/domain/user"

type CreateUserRepositoryInterface interface {
	Execute(user *user.User) (*user.User, error)
}

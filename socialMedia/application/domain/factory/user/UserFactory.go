package factory

import (
	"socialMedia/application/domain/user"
	valueObjectEmail "socialMedia/application/domain/valueObject/email"
	valueObjectName "socialMedia/application/domain/valueObject/name"
	valueObjectNickname "socialMedia/application/domain/valueObject/nickname"
	valueObjectPassword "socialMedia/application/domain/valueObject/password"
	valueObjectUuid "socialMedia/application/domain/valueObject/uuid"
)

type UserFactory struct{}

func (userFactory *UserFactory) MakeComplete(
	id valueObjectUuid.Uuid,
	name valueObjectName.Name,
	nickname valueObjectNickname.Nickname,
	email valueObjectEmail.Email,
	password valueObjectPassword.Password,
) *user.User {
	return user.NewUser(id, name, nickname, email, password)
}

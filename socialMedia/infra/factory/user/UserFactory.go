package factory

import (
	"socialMedia/application/domain/user"
	valueObjectEmail "socialMedia/application/domain/valueObject/email"
	valueObjectName "socialMedia/application/domain/valueObject/name"
	valueObjectNickname "socialMedia/application/domain/valueObject/nickname"
	valueObjectPassword "socialMedia/application/domain/valueObject/password"
	valueObjectUuid "socialMedia/application/domain/valueObject/uuid"
	"socialMedia/infra/uuid"

	infraEmail "socialMedia/infra/email"
	infraPassword "socialMedia/infra/password"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (userFactory *UserFactory) MakeComplete(
	name string,
	nickname string,
	email string,
	password string,
) (*user.User, error) {
	uuidValueObject := valueObjectUuid.NewUuid(uuid.NewUuidGenerator())

	nameValueObject, err := valueObjectName.NewName(name)

	if err != nil {
		return nil, err
	}

	nicknameValueObject, err := valueObjectNickname.NewNickname(nickname)

	if err != nil {
		return nil, err
	}

	emailValueObject, err := valueObjectEmail.NewEmail(email, infraEmail.NewEmailValidator())

	if err != nil {
		return nil, err
	}

	passwordValueObject, err := valueObjectPassword.NewPassword(
		password,
		infraPassword.NewPasswordValidator(),
		infraPassword.NewPasswordEncrypter(),
	)

	if err != nil {
		return nil, err
	}

	return user.NewUser(
		*uuidValueObject,
		*nameValueObject,
		*nicknameValueObject,
		*emailValueObject,
		*passwordValueObject,
	), nil
}

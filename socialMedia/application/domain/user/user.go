package user

import (
	valueObjectEmail "socialMedia/application/domain/valueObject/email"
	valueObjectName "socialMedia/application/domain/valueObject/name"
	valueObjectNickname "socialMedia/application/domain/valueObject/nickname"
	valueObjectPassword "socialMedia/application/domain/valueObject/password"
	valueObjectUuid "socialMedia/application/domain/valueObject/uuid"
)

type User struct {
	id       valueObjectUuid.Uuid
	name     valueObjectName.Name
	nickname valueObjectNickname.Nickname
	email    valueObjectEmail.Email
	password valueObjectPassword.Password
}

func NewUser(
	id valueObjectUuid.Uuid,
	name valueObjectName.Name,
	nickname valueObjectNickname.Nickname,
	email valueObjectEmail.Email,
	password valueObjectPassword.Password,
) *User {
	return &User{
		id:       id,
		name:     name,
		nickname: nickname,
		email:    email,
		password: password,
	}
}

func (user *User) GetId() string {
	return user.id.GetValue()
}

func (user *User) GetName() string {
	return user.name.GetValue()
}

func (user *User) GetNickname() string {
	return user.nickname.GetValue()
}

func (user *User) GetEmail() string {
	return user.email.GetValue()
}

func (user *User) GetPassword() valueObjectPassword.Password {
	return user.password
}

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

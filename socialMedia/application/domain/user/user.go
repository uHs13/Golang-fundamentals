package user

import (
	valueObjectEmail "socialMedia/application/domain/valueObject/email"
	valueObjectName "socialMedia/application/domain/valueObject/name"
	valueObjectNickname "socialMedia/application/domain/valueObject/nickname"
	valueObjectPassword "socialMedia/application/domain/valueObject/password"
	valueObjectUuid "socialMedia/application/domain/valueObject/uuid"
)

type User struct {
	Id       valueObjectUuid.Uuid
	Name     valueObjectName.Name
	Nickname valueObjectNickname.Nickname
	Email    valueObjectEmail.Email
	Password valueObjectPassword.Password
}

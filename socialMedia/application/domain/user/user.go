package user

import (
	valueObjectName "socialMedia/application/domain/valueObject/name"
	valueObjectUuid "socialMedia/application/domain/valueObject/uuid"
)

type User struct {
	Id       valueObjectUuid.Uuid
	Name     valueObjectName.Name
	Nickname string
	Email    string
	Password string
}

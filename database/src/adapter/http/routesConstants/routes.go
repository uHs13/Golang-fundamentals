package routesConstants

import "net/http"

const (
	PostCreateUser = "/user"
	GetUsers       = "/users"
	GetUser        = "/user/{id}"

	UsersKeyConst = "users"
	UserKeyConst  = "user"

	MessageKeyConst    = "message"
	MessageUserCreated = "User successfully created!"

	BadRequestConst     = http.StatusBadRequest
	CreatedRequestConst = http.StatusCreated
)

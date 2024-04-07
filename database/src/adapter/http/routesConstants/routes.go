package routesConstants

import "net/http"

const (
	PostCreateUser = "/user"
	GetUsers       = "/users"

	MessageKeyConst    = "message"
	MessageUserCreated = "User successfully created!"

	BadRequestConst     = http.StatusBadRequest
	CreatedRequestConst = http.StatusCreated
)

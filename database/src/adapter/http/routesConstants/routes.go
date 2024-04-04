package routesConstants

import "net/http"

const (
	PostCreateUser = "/user"

	MessageKeyConst    = "message"
	MessageUserCreated = "User successfully created!"

	BadRequestConst     = http.StatusBadRequest
	CreatedRequestConst = http.StatusCreated
)

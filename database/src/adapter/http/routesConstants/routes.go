package routesConstants

import "net/http"

const (
	PostCreateUser = "/user"
	GetUsers       = "/users"
	GetUser        = "/user/{id}"
	PatchUser      = "/user/{id}"
	DeleteUser     = "/user/{id}"

	UsersKeyConst = "users"
	UserKeyConst  = "user"

	MessageKeyConst     = "message"
	MessageUserCreated  = "user successfully created!"
	MessageUserNotFound = "user not found"
	MessageUserDeleted  = "user successfully deleted!"

	BadRequestConst     = http.StatusBadRequest
	CreatedRequestConst = http.StatusCreated
	OkRequestConst      = http.StatusOK
)

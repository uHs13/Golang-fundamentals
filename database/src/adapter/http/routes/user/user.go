package user

import (
	"database/sql"
	"database/src/adapter/http/handlers/userHandler"
	"database/src/adapter/http/routesConstants"
	"database/src/core/port"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PostCreateHandlerKeyConst string = "postCreateUserHandler"
	GetUsersHandlerKeyConst   string = "getUsersHandler"
	GetUserHandlerKeyConst    string = "getUserHandler"
	PatchUserHandlerKeyConst  string = "patchUserHandler"
	DeleteUserHandlerKeyConst string = "deleteUserHandler"
)

type UserRoutes struct {
	app          *mux.Router
	userHandlers map[string]port.HandlerInterface
}

func NewUserRoutes(
	app *mux.Router,
	connection *sql.DB,
) port.RoutesInterface {
	return &UserRoutes{
		app,
		createUserHandlers(connection),
	}
}

func (userRoutes *UserRoutes) Register() {
	userRoutes.app.HandleFunc(
		routesConstants.PostCreateUser,
		userRoutes.userHandlers[PostCreateHandlerKeyConst].Handle,
	).Methods(http.MethodPost)

	userRoutes.app.HandleFunc(
		routesConstants.GetUsers,
		userRoutes.userHandlers[GetUsersHandlerKeyConst].Handle,
	).Methods(http.MethodGet)

	userRoutes.app.HandleFunc(
		routesConstants.GetUser,
		userRoutes.userHandlers[GetUserHandlerKeyConst].Handle,
	).Methods(http.MethodGet)

	userRoutes.app.HandleFunc(
		routesConstants.PatchUser,
		userRoutes.userHandlers[PatchUserHandlerKeyConst].Handle,
	).Methods(http.MethodPatch)

	userRoutes.app.HandleFunc(
		routesConstants.PatchUser,
		userRoutes.userHandlers[DeleteUserHandlerKeyConst].Handle,
	).Methods(http.MethodDelete)
}

func createUserHandlers(connection *sql.DB) map[string]port.HandlerInterface {
	return map[string]port.HandlerInterface{
		PostCreateHandlerKeyConst: userHandler.NewCreateUserHandler(connection),
		GetUsersHandlerKeyConst:   userHandler.NewFindUsersHandler(connection),
		GetUserHandlerKeyConst:    userHandler.NewFindUserHandler(connection),
		PatchUserHandlerKeyConst:  userHandler.NewEditUserHandler(connection),
		DeleteUserHandlerKeyConst: userHandler.NewDeleteUserHandler(connection),
	}
}

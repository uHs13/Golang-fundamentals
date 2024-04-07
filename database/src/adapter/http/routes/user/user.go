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
}

func createUserHandlers(connection *sql.DB) map[string]port.HandlerInterface {
	return map[string]port.HandlerInterface{
		PostCreateHandlerKeyConst: userHandler.NewCreateUserHandler(connection),
		GetUsersHandlerKeyConst:   userHandler.NewFindUsersHandler(connection),
	}
}

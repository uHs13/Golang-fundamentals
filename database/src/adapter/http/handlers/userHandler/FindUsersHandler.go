package userHandler

import (
	"database/sql"
	"database/src/adapter/http/routes"
	"database/src/adapter/http/routesConstants"
	"database/src/core/port"
	"database/src/core/port/repositories"
	findUsers "database/src/core/useCase/findUsers"
	"database/src/infra/database/repository"
	"net/http"
)

type FindUsersHandler struct {
	Connection   *sql.DB
	UserDatabase repositories.UserRepository
}

func NewFindUsersHandler(connection *sql.DB) port.HandlerInterface {
	return &FindUsersHandler{
		Connection: connection,
	}
}

func (findUsersHandler *FindUsersHandler) Handle(
	w http.ResponseWriter,
	r *http.Request,
) {
	users, err := findUsers.NewFindUsers(
		repository.NewUserDatabase(findUsersHandler.Connection),
	).Execute()

	if err != nil {
		routes.NewJsonResponse(
			w,
			err.Error(),
			routesConstants.BadRequestConst,
		).SendSimpleJson()

		return
	}

	routes.NewJsonResponse(
		w,
		"",
		routesConstants.OkRequestConst,
	).SendArrayJson(routesConstants.UsersKeyConst, users)
}

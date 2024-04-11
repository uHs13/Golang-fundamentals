package userHandler

import (
	"database/sql"
	"database/src/adapter/http/routes"
	"database/src/adapter/http/routesConstants"
	userDomain "database/src/core/domain/user"
	"database/src/core/port"
	"database/src/core/port/repositories"
	"database/src/core/useCase/findUser"
	"database/src/infra/database/repository"
	"database/src/infra/requestEntity/userRequestEntity"
	"net/http"
)

type FindUserHandler struct {
	Connection   *sql.DB
	UserDatabase repositories.UserRepository
}

func NewFindUserHandler(connection *sql.DB) port.HandlerInterface {
	return &FindUserHandler{
		Connection: connection,
	}
}

func (findUserHandler *FindUserHandler) Handle(
	w http.ResponseWriter,
	r *http.Request,
) {
	user, err := findUserHandler.defineFindUser(r)

	if err != nil {
		routes.NewJsonResponse(
			w,
			err.Error(),
			routesConstants.BadRequestConst,
		).SendSimpleJson()

		return
	}

	foundUser, err := findUser.NewFindUser(
		repository.NewUserDatabase(findUserHandler.Connection),
		*user,
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
	).SendArrayJson(routesConstants.UserKeyConst, foundUser)
}

func (findUserHandler *FindUserHandler) defineFindUser(
	r *http.Request,
) (*userDomain.User, error) {
	userRequest, err := userRequestEntity.DecodeFindUser(r)

	if err != nil {
		return nil, err
	}

	if err := userRequest.ValidateId(); err != nil {
		return nil, err
	}

	user := userDomain.NewUser(
		userRequest.Id,
		"",
		"",
		"",
		"",
		"",
	)

	return user, nil
}

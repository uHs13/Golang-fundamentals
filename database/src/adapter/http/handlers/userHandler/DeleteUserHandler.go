package userHandler

import (
	"database/sql"
	"database/src/adapter/http/routes"
	"database/src/adapter/http/routesConstants"
	datetime "database/src/core/domain/dateTime"
	userDomain "database/src/core/domain/user"
	"database/src/core/port"
	"database/src/core/port/repositories"
	"database/src/core/useCase/deleteUser"
	"database/src/infra/database/repository"
	"database/src/infra/requestEntity/userRequestEntity"
	"net/http"
)

type DeleteUserHandler struct {
	Connection   *sql.DB
	UserDatabase repositories.UserRepository
}

func NewDeleteUserHandler(connection *sql.DB) port.HandlerInterface {
	return &DeleteUserHandler{
		Connection: connection,
	}
}

func (deleteUserHandler *DeleteUserHandler) Handle(
	w http.ResponseWriter,
	r *http.Request,
) {
	user, err := deleteUserHandler.defineDeleteUser(r)

	if err != nil {
		routes.NewJsonResponse(
			w,
			err.Error(),
			routesConstants.BadRequestConst,
		).SendSimpleJson()

		return
	}

	err = deleteUser.NewDeleteUser(
		repository.NewUserDatabase(deleteUserHandler.Connection),
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
		routesConstants.MessageUserDeleted,
		routesConstants.OkRequestConst,
	).SendSimpleJson()
}

func (deleteUserHandler *DeleteUserHandler) defineDeleteUser(
	r *http.Request,
) (*userDomain.User, error) {
	userRequest, err := userRequestEntity.DecodeDeleteUser(r)

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
		datetime.BuildFormattedCurrentDate(),
	)

	return user, nil
}

package userHandler

import (
	"database/sql"
	"database/src/adapter/http/routes"
	"database/src/adapter/http/routesConstants"
	datetime "database/src/core/domain/dateTime"
	userDomain "database/src/core/domain/user"
	"database/src/core/port"
	"database/src/core/port/repositories"
	"database/src/core/useCase/editUser"
	"database/src/infra/database/repository"
	"database/src/infra/requestEntity/userRequestEntity"
	"net/http"
)

type EditUserHandler struct {
	Connection   *sql.DB
	UserDatabase repositories.UserRepository
}

func NewEditUserHandler(connection *sql.DB) port.HandlerInterface {
	return &EditUserHandler{
		Connection: connection,
	}
}

func (editUserHandler *EditUserHandler) Handle(
	w http.ResponseWriter,
	r *http.Request,
) {
	user, err := editUserHandler.defineEditUser(r)

	if err != nil {
		routes.NewJsonResponse(
			w,
			err.Error(),
			routesConstants.BadRequestConst,
		).SendSimpleJson()

		return
	}

	err = editUser.NewEditUser(
		repository.NewUserDatabase(editUserHandler.Connection),
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
	).SendArrayJson(routesConstants.UserKeyConst, user)
}

func (editUserHandler *EditUserHandler) defineEditUser(
	r *http.Request,
) (*userDomain.User, error) {
	userRequest, err := userRequestEntity.DecodeEditUser(r)

	if err != nil {
		return nil, err
	}

	if err = userRequest.ValidateEdit(); err != nil {
		return nil, err
	}

	user := userDomain.NewUser(
		userRequest.Id,
		userRequest.Name,
		userRequest.Email,
		"",
		datetime.BuildFormattedCurrentDate(),
		"",
	)

	return user, nil
}

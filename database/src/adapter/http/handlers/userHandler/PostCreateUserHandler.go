package userHandler

import (
	"database/sql"
	"database/src/adapter/http/routes"
	"database/src/adapter/http/routesConstants"
	datetime "database/src/core/domain/dateTime"
	userDomain "database/src/core/domain/user"
	"database/src/core/port"
	"database/src/core/port/repositories"
	createUser "database/src/core/useCase/createUser"
	"database/src/infra/database/repository"
	"database/src/infra/requestEntity/userRequestEntity"
	"net/http"

	"github.com/google/uuid"
)

type CreateUserHandler struct {
	Connection   *sql.DB
	UserDatabase repositories.UserRepository
}

func NewCreateUserHandler(connection *sql.DB) port.HandlerInterface {
	return &CreateUserHandler{
		Connection: connection,
	}
}

func (createUserHandler *CreateUserHandler) Handle(
	w http.ResponseWriter,
	r *http.Request,
) {
	user, err := createUserHandler.defineCreateUser(r)

	if err != nil {
		routes.NewJsonResponse(
			w,
			err.Error(),
			routesConstants.BadRequestConst,
		).SendSimpleJson()

		return
	}

	err = createUser.NewCreateUser(
		repository.NewUserDatabase(createUserHandler.Connection),
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
		routesConstants.MessageUserCreated,
		routesConstants.CreatedRequestConst,
	).SendSimpleJson()
}

func (createUserHandler *CreateUserHandler) defineCreateUser(
	r *http.Request,
) (*userDomain.User, error) {
	userRequest, err := userRequestEntity.DecodeCreateUserRequest(r)

	if err != nil {
		return nil, err
	}

	if err := userRequest.Validate(); err != nil {
		return nil, err
	}

	user := userDomain.NewUser(
		uuid.NewString(),
		userRequest.Name,
		userRequest.Email,
		datetime.BuildFormattedCurrentDate(),
	)

	return user, nil
}

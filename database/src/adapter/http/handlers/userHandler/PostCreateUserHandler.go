package userHandler

import (
	"database/sql"
	"database/src/adapter/http/routes"
	"database/src/adapter/http/routesConstants"
	datetime "database/src/core/domain/dateTime"
	userDomain "database/src/core/domain/user"
	"database/src/core/port"
	"database/src/core/port/repositories"
	userUseCase "database/src/core/useCase"
	"database/src/infra/database/repository"
	"database/src/infra/requestEntity/userRequestEntity"
	"fmt"
	"net/http"
	"os"

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
			routesConstants.MessageKeyConst,
			err.Error(),
			routesConstants.BadRequestConst,
		).ThrowError()

		return
	}

	err = userUseCase.NewCreateUser(
		repository.NewUserDatabase(createUserHandler.Connection),
		*user,
	).Execute()

	if err != nil {
		routes.NewJsonResponse(
			w,
			routesConstants.MessageKeyConst,
			err.Error(),
			routesConstants.BadRequestConst,
		).ThrowError()

		return
	}

	fmt.Println(user)
	os.Exit(1)
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

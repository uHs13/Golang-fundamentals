package userHandler

import (
	"database/sql"
	"database/src/core/port"
	"fmt"
	"net/http"
)

type CreateUserHandler struct {
	Connection *sql.DB
}

func NewCreateUserHandler(connection *sql.DB) port.HandlerInterface {
	return &CreateUserHandler{
		Connection: connection,
	}
}

func (createUserHandler *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("That is it!")
}

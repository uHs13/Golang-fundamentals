package userRequestEntity

import (
	"database/src/infra/requestEntity/requestEntityFieldsValidation"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteUserRequest struct {
	Id string `json:"id"`
}

func DecodeDeleteUser(r *http.Request) (*DeleteUserRequest, error) {
	params := mux.Vars(r)[Id]

	if params == "" {
		return nil, fmt.Errorf(EmptyIdErrorMsg)
	}

	user := &DeleteUserRequest{
		Id: params,
	}

	return user, nil
}

func (deleteUserRequest *DeleteUserRequest) ValidateId() error {
	if err := requestEntityFieldsValidation.ValidateId(deleteUserRequest.Id); err != nil {
		return err
	}

	return nil
}

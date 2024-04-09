package userRequestEntity

import (
	"database/src/infra/requestEntity/requestEntityFieldsValidation"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type FindUserRequest struct {
	Id string `json:"id"`
}

const (
	EmptyIdErrorMsg = "the id must not be empty"
	Id              = "id"
)

func DecodeFindUser(r *http.Request) (*FindUserRequest, error) {
	params := mux.Vars(r)[Id]

	if params == "" {
		return nil, fmt.Errorf(EmptyIdErrorMsg)
	}

	user := &FindUserRequest{
		Id: params,
	}

	return user, nil
}

func (findUserRequest *FindUserRequest) ValidateId() error {
	if err := requestEntityFieldsValidation.ValidateId(findUserRequest.Id); err != nil {
		return err
	}

	return nil
}

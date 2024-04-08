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
	emptyIdErrorMsg = "the id must not be empty"
)

func DecodeFindUser(r *http.Request) (*FindUserRequest, error) {
	params := mux.Vars(r)["id"]

	if params == "" {
		return nil, fmt.Errorf(emptyIdErrorMsg)
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

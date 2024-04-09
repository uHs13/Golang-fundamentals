package userRequestEntity

import (
	"database/src/infra/requestEntity/requestEntityFieldsValidation"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type EditUserRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func DecodeEditUser(r *http.Request) (*EditUserRequest, error) {
	params := mux.Vars(r)[Id]

	if params == "" {
		return nil, fmt.Errorf(EmptyIdErrorMsg)
	}

	var user *EditUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}

	user.Id = params

	return user, nil
}

func (editUserRequest *EditUserRequest) ValidateEdit() error {
	if err := requestEntityFieldsValidation.ValidateId(editUserRequest.Id); err != nil {
		return err
	}

	if err := requestEntityFieldsValidation.ValidateName(editUserRequest.Name); err != nil {
		return err
	}

	if err := requestEntityFieldsValidation.ValidateEmail(editUserRequest.Email); err != nil {
		return err
	}

	return nil
}

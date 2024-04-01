package userRequestEntity

import (
	"database/src/infra/requestEntity/requestEntityFieldsValidation"
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func DecodeCreateUserRequest(r *http.Request) (*CreateUserRequest, error) {
	var user *CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *CreateUserRequest) Validate() error {
	if err := requestEntityFieldsValidation.ValidateName(user.Name); err != nil {
		return err
	}

	if err := requestEntityFieldsValidation.ValidateEmail(user.Email); err != nil {
		return err
	}

	return nil
}

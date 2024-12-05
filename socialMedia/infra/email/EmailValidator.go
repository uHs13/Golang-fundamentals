package email

import (
	"errors"
	"fmt"
	"net/mail"
)

type EmailValidator struct{}

func NewEmailValidator() *EmailValidator {
	return &EmailValidator{}
}

func (emailValidator *EmailValidator) Validate(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		errorMessage := fmt.Sprintf("the provided email '%s' is not valid", email)
		return errors.New(errorMessage)
	}

	return nil
}

package requestEntityFieldsValidation

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

const (
	maxNameLength          = 100
	nameMaxLengthErrorMsg  = "the name cannot be greater than 100 characters"
	nameMinLengthErrorMsg  = "the name cannot be empty"
	maxEmailLength         = 100
	emailMaxLengthErrorMsg = "the email cannot be greater than 100 characters"
	emailMinLengthErrorMsg = "the email cannot be empty"
	invalidEmailErrorMsg   = "the provided email is not valid"
)

func ValidateName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf(nameMinLengthErrorMsg)
	}

	if len(name) > maxNameLength {
		return fmt.Errorf(nameMaxLengthErrorMsg)
	}

	return nil
}

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return fmt.Errorf(emailMaxLengthErrorMsg)
	}

	if len(email) > maxEmailLength {
		return fmt.Errorf(emailMaxLengthErrorMsg)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return fmt.Errorf(invalidEmailErrorMsg)
	}

	return nil
}

func ValidateId(hash string) error {
	if _, err := uuid.Parse(hash); err != nil {
		return err
	}

	return nil
}

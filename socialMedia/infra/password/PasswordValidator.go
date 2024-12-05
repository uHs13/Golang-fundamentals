package password

import (
	"errors"
	"fmt"
	"unicode"
)

const (
	passwordMinLengthConst = 8
)

type PasswordValidator struct {
	plainPassword string
	hasUpperCase  bool
	hasLowerCase  bool
	hasNumber     bool
}

func NewPasswordValidator() *PasswordValidator {
	return &PasswordValidator{}
}

func (passwordValidator *PasswordValidator) Validate(plainPassword string) error {
	passwordValidator.plainPassword = plainPassword

	if err := passwordValidator.checkLength(); err != nil {
		return err
	}

	if err := passwordValidator.executeDeeperValidations(); err != nil {
		return err
	}

	return nil
}

func (passwordValidator *PasswordValidator) checkLength() error {
	if len(passwordValidator.plainPassword) < passwordMinLengthConst {
		errorMessage := fmt.Sprintf("the password must have at least '%d' characters", passwordMinLengthConst)

		return errors.New(errorMessage)
	}

	return nil
}

func (passwordValidator *PasswordValidator) executeDeeperValidations() error {
	for _, char := range passwordValidator.plainPassword {
		if unicode.IsUpper(char) {
			passwordValidator.hasUpperCase = true
		}

		if unicode.IsLower(char) {
			passwordValidator.hasLowerCase = true
		}

		if unicode.IsDigit(char) {
			passwordValidator.hasNumber = true
		}
	}

	return passwordValidator.checkForRuleErrors()
}

func (passwordValidator *PasswordValidator) checkForRuleErrors() error {
	if !passwordValidator.hasUpperCase {
		return errors.New("the password must contain at least one upper case letter")
	}

	if !passwordValidator.hasLowerCase {
		return errors.New("the password must contain at least one lower case letter")
	}

	if !passwordValidator.hasNumber {
		return errors.New("the password must contain at least one number")
	}

	return nil
}

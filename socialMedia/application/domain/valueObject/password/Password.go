package valueObject

import port "socialMedia/application/port/password"

type Password struct {
	plainValue string
}

func NewPassword(
	plainValue string,
	validator port.PasswordValidatorInterface,
) (*Password, error) {
	if err := validator.Validate(plainValue); err != nil {
		return nil, err
	}

	return &Password{plainValue: plainValue}, nil
}

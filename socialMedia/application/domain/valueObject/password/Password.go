package valueObject

import port "socialMedia/application/port/password"

type Password struct {
	plainValue        string
	passwordEncrypter port.PasswordEncrypterInterface
}

func NewPassword(
	plainValue string,
	validator port.PasswordValidatorInterface,
	passwordEncrypter port.PasswordEncrypterInterface,
) (*Password, error) {
	if err := validator.Validate(plainValue); err != nil {
		return nil, err
	}

	return &Password{
		plainValue:        plainValue,
		passwordEncrypter: passwordEncrypter,
	}, nil
}

func (password *Password) GenerateHash() (string, error) {
	hash, err := password.passwordEncrypter.GenerateHash(password.plainValue)

	if err != nil {
		return "", err
	}

	return hash, nil
}

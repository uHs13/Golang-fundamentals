package valueObject

import port "socialMedia/application/port/email"

type Email struct {
	value string
}

func NewEmail(email string, validator port.EmailValidatorInterface) (*Email, error) {
	if err := validator.Validate(email); err != nil {
		return nil, err
	}

	valueObject := &Email{value: email}

	return valueObject, nil
}

func (email *Email) GetValue() string {
	return email.value
}

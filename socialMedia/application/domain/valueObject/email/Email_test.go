package valueObject_test

import (
	"fmt"
	valueObject "socialMedia/application/domain/valueObject/email"
	"socialMedia/infra/email"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyCreateAnEmail(t *testing.T) {
	validEmail := "john.cena@email.com"
	emailValueObject, err := valueObject.NewEmail(validEmail, email.NewEmailValidator())

	require.Nil(t, err)
	require.Equal(t, validEmail, emailValueObject.GetValue())
}

func TestShouldThrowAnErrorWhenEmailIsNotValid(t *testing.T) {
	invalidEmail := "john.cena"
	_, err := valueObject.NewEmail(invalidEmail, email.NewEmailValidator())

	require.Equal(t, err.Error(), fmt.Sprintf("the provided email '%s' is not valid", invalidEmail))
}

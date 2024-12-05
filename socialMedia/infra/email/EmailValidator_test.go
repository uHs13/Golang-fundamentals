package email_test

import (
	"fmt"
	"socialMedia/infra/email"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyValidateAnEmail(t *testing.T) {
	emailValidator := email.NewEmailValidator()

	err := emailValidator.Validate("some@email.com")

	require.Nil(t, err)
}

func TestShouldThrowAnErrorWhenEmailIsNotValid(t *testing.T) {
	emailValidator := email.NewEmailValidator()

	invalidEmail := "someemail.com"

	err := emailValidator.Validate(invalidEmail)

	errorMessage := fmt.Sprintf("the provided email '%s' is not valid", invalidEmail)

	require.Equal(t, err.Error(), errorMessage)
}

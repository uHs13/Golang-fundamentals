package password_test

import (
	"socialMedia/infra/password"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldThrowAnErrorWhenPasswordDoesNotMeetTheMinimumLengthRequirements(t *testing.T) {
	validator := password.NewPasswordValidator()

	err := validator.Validate("short")

	require.Equal(t, err.Error(), "the password must have at least '8' characters")
}

func TestShouldThrowAnErrorWhenPasswordDoesNotContainsAtLeastOneUpperCaseLetter(t *testing.T) {
	validator := password.NewPasswordValidator()

	err := validator.Validate("passwordwithnouppercaseletter")

	require.Equal(t, err.Error(), "the password must contain at least one upper case letter")
}

func TestShouldThrowAnErrorWhenPasswordDoesNotContainsAtLeastOneLowerCaseLetter(t *testing.T) {
	validator := password.NewPasswordValidator()

	err := validator.Validate("ONLYUPPERCASE")

	require.Equal(t, err.Error(), "the password must contain at least one lower case letter")
}

func TestShouldThrowAnErrorWhenPasswordDoesNotContainsAtLeastOneNumber(t *testing.T) {
	validator := password.NewPasswordValidator()

	err := validator.Validate("passwordwithnoNumber")

	require.Equal(t, err.Error(), "the password must contain at least one number")
}

func TestMustProperlyCreateAPassword(t *testing.T) {
	validator := password.NewPasswordValidator()

	err := validator.Validate("AeR2@rty")

	require.Nil(t, err)
}

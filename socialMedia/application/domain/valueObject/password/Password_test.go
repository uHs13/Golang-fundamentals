package valueObject_test

import (
	valueObject "socialMedia/application/domain/valueObject/password"
	"socialMedia/infra/password"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyCreateAPassword(t *testing.T) {
	_, err := valueObject.NewPassword(
		"My@P4ssW0rd",
		password.NewPasswordValidator(),
		password.NewPasswordEncrypter(),
	)

	require.Nil(t, err)
}

func TestShouldThrowAnErrorWhenPasswordDoesNotMeetTheMinimumLengthRequirements(t *testing.T) {
	_, err := valueObject.NewPassword(
		"short",
		password.NewPasswordValidator(),
		password.NewPasswordEncrypter(),
	)

	require.Equal(t, err.Error(), "the password must have at least '8' characters")
}

func TestShouldThrowAnErrorWhenPasswordDoesNotContainsAtLeastOneUpperCaseLetter(t *testing.T) {
	_, err := valueObject.NewPassword(
		"passwordwithnouppercaseletter",
		password.NewPasswordValidator(),
		password.NewPasswordEncrypter(),
	)

	require.Equal(t, err.Error(), "the password must contain at least one upper case letter")
}

func TestShouldThrowAnErrorWhenPasswordDoesNotContainsAtLeastOneLowerCaseLetter(t *testing.T) {
	_, err := valueObject.NewPassword(
		"ONLYUPPERCASE",
		password.NewPasswordValidator(),
		password.NewPasswordEncrypter(),
	)

	require.Equal(t, err.Error(), "the password must contain at least one lower case letter")
}

func TestShouldThrowAnErrorWhenPasswordDoesNotContainsAtLeastOneNumber(t *testing.T) {
	_, err := valueObject.NewPassword(
		"passwordwithnoNumber",
		password.NewPasswordValidator(),
		password.NewPasswordEncrypter(),
	)

	require.Equal(t, err.Error(), "the password must contain at least one number")
}

func TestShouldProperlyGeneratePasswordHash(t *testing.T) {
	password, err := valueObject.NewPassword(
		"My@P4ssW0rd",
		password.NewPasswordValidator(),
		password.NewPasswordEncrypter(),
	)

	require.Nil(t, err)

	hash, err := password.GenerateHash()

	require.Nil(t, err)
	require.NotEmpty(t, hash)
}

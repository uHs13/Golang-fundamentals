package password_test

import (
	"socialMedia/infra/password"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyGeneratePasswordHash(t *testing.T) {
	plainText := "someText"

	hash, err := password.NewPasswordEncrypter().GenerateHash(plainText)

	require.Nil(t, err)
	require.NotEmpty(t, hash)
}

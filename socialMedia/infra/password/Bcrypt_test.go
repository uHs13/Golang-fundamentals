package password_test

import (
	"socialMedia/infra/password"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyGenerateHashFromPlainText(t *testing.T) {
	bcrypt := password.NewBcrypt()

	hash, err := bcrypt.Encrypt("plainText")

	require.Nil(t, err)
	require.NotEmpty(t, hash)
}

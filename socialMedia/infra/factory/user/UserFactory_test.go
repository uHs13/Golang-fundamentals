package factory_test

import (
	factory "socialMedia/infra/factory/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyMakeACompleteUser(t *testing.T) {
	name := "John Cena"
	nickname := "john.cena"
	email := "john.cena@email.com"
	password := "John@12345*"

	user, err := factory.NewUserFactory().MakeComplete(
		name,
		nickname,
		email,
		password,
	)

	require.Nil(t, err)

	require.NotEmpty(t, user.GetId())
	require.Equal(t, user.GetName(), name)
	require.Equal(t, user.GetNickname(), nickname)
	require.Equal(t, user.GetEmail(), email)
}

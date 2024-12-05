package valueObject_test

import (
	valueObject "socialMedia/application/domain/valueObject/nickname"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyCreateANickname(t *testing.T) {
	nickName := "johnCena13"
	nickNamevalueObject, err := valueObject.NewNickname(nickName)

	require.Nil(t, err)
	require.Equal(t, nickNamevalueObject.GetValue(), nickName)
}

func TestShouldThrowAnErrorWhenNicknameIsTooLong(t *testing.T) {
	_, err := valueObject.NewNickname("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident")

	require.Equal(t, err.Error(), "the nickname length must not exceed '255' characters")
}

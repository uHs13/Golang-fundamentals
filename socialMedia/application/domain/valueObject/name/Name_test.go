package valueObject_test

import (
	valueObject "socialMedia/application/domain/valueObject/name"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyCreateAName(t *testing.T) {
	name := "John Cena"
	nameValueObject, err := valueObject.NewName(name)

	require.Nil(t, err)
	require.Equal(t, nameValueObject.GetValue(), name)
}

func TestShouldThrowAnErrorWhenNameIsEmpty(t *testing.T) {
	_, err := valueObject.NewName("")

	require.Equal(t, err.Error(), "the name cannot be empty")
}

func TestShouldThrowAnErrorWhenNameIsNotComposedByAtLeastTwoWords(t *testing.T) {
	_, err := valueObject.NewName("John")

	require.Equal(t, err.Error(), "the name must contain at least two words")
}

func TestShouldThrowAnErrorWhenNameIsTooLong(t *testing.T) {
	_, err := valueObject.NewName("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia")

	require.Equal(t, err.Error(), "the name cannot be longer than '255' characters")
}

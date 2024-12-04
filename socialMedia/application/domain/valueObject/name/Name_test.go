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

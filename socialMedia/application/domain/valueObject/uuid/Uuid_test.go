package valueObject_test

import (
	valueObject "socialMedia/application/domain/valueObject/uuid"
	"socialMedia/infra/uuid"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyCreateAnUuid(t *testing.T) {
	uuidGenerator := uuid.NewUuidGenerator()

	createdUuid := valueObject.NewUuid(uuidGenerator)

	require.NotEmpty(t, createdUuid.GetValue())
}

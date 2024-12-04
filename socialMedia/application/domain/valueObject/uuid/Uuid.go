package valueObject

import port "socialMedia/application/port/uuid"

type Uuid struct {
	value string
}

func NewUuid(uuidGenerator port.UuidGeneratorInterface) *Uuid {
	return &Uuid{
		value: uuidGenerator.Generate(),
	}
}

func (uuid *Uuid) GetValue() string {
	return uuid.value
}

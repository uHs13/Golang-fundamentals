package valueObject

import (
	"errors"
	"strings"
)

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	valueObject := &Name{value: name}

	if err := valueObject.validate(); err != nil {
		return nil, err
	}

	return valueObject, nil
}

func (name *Name) GetValue() string {
	return name.value
}

func (name *Name) validate() error {
	if len(name.GetValue()) == 0 {
		return errors.New("the name cannot be empty")
	}

	nameComponents := strings.Split(name.GetValue(), " ")

	if len(nameComponents) < 2 {
		return errors.New("the name must contain at least two words")
	}

	return nil
}

package valueObject

import (
	"errors"
	"fmt"
	"strings"
)

const (
	nameMaxLengthConst          = 255
	nameMinimumWordsAmountConst = 2
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
	if err := name.validateNameLength(); err != nil {
		return err
	}

	if err := name.validateNameMinimumWordsAmount(); err != nil {
		return err
	}

	return nil
}

func (name *Name) validateNameLength() error {
	if len(name.GetValue()) == 0 {
		return errors.New("the name cannot be empty")
	}

	if len(name.GetValue()) > nameMaxLengthConst {
		return fmt.Errorf("the name cannot be longer than '%d' characters", nameMaxLengthConst)
	}

	return nil
}

func (name *Name) validateNameMinimumWordsAmount() error {
	nameComponents := strings.Split(name.GetValue(), " ")

	if len(nameComponents) < nameMinimumWordsAmountConst {
		return errors.New("the name must contain at least two words")
	}

	return nil
}

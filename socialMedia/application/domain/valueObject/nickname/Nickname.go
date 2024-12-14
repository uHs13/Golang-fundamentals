package valueObject

import (
	"errors"
	"fmt"
)

const (
	nicknameMaxLengthConst = 255
)

type Nickname struct {
	value string
}

func NewNickname(nickName string) (*Nickname, error) {
	valueObject := &Nickname{value: nickName}

	if err := valueObject.validate(); err != nil {
		return nil, err
	}

	return valueObject, nil
}

func (nickname *Nickname) validate() error {
	if len(nickname.GetValue()) == 0 {
		return errors.New("the nickname cannot be empty")
	}

	if len(nickname.GetValue()) > nicknameMaxLengthConst {
		errorMessage := fmt.Sprintf("the nickname length must not exceed '%d' characters", nicknameMaxLengthConst)
		return errors.New(errorMessage)
	}

	return nil
}

func (nickname *Nickname) GetValue() string {
	return nickname.value
}

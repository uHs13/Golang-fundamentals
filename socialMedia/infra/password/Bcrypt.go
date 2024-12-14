package password

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (bc *Bcrypt) Encrypt(plainText string) (string, error) {
	bytePassword := []byte(plainText)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

package password

type PasswordEncrypter struct{}

func NewPasswordEncrypter() *PasswordEncrypter {
	return &PasswordEncrypter{}
}

func (encrypter *PasswordEncrypter) GenerateHash(plainText string) (string, error) {
	encryptMechanism := encrypter.generateEncryptMechanism()

	hash, err := encryptMechanism.Encrypt(plainText)

	if err != nil {
		return "", err
	}

	return hash, nil
}

func (encrypter *PasswordEncrypter) generateEncryptMechanism() EncryptMechanismInterface {
	return NewEncryptMechanismFactory().Make()
}

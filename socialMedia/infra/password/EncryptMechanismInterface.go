package password

type EncryptMechanismInterface interface {
	Encrypt(plainText string) (string, error)
}

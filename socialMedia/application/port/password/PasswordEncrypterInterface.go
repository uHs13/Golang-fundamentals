package port

type PasswordEncrypterInterface interface {
	GenerateHash(passwordPlainText string) (string, error)
}

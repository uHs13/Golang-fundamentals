package port

type PasswordValidatorInterface interface {
	Validate(plainPassword string) error
}

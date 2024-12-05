package port

type EmailValidatorInterface interface {
	Validate(email string) error
}

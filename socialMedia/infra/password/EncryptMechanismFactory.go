package password

type EncryptMechanismFactory struct {
}

func NewEncryptMechanismFactory() *EncryptMechanismFactory {
	return &EncryptMechanismFactory{}
}

func (factory *EncryptMechanismFactory) Make() EncryptMechanismInterface {
	return NewBcrypt()
}

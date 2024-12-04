package port

type DatabaseFactoryInterface interface {
	MakeInstance() (DatabaseConnectionInterface, error)
}

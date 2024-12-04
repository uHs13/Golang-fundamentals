package database

import (
	"errors"
	port "socialMedia/application/port/database"
)

type DatabaseFactory struct {
	ConnectionType string
}

func NewDatabaseFactory(connectionType string) *DatabaseFactory {
	return &DatabaseFactory{
		ConnectionType: connectionType,
	}
}

func (databaseFactory *DatabaseFactory) MakeInstance() (
	port.DatabaseConnectionInterface,
	error,
) {
	if databaseFactory.ConnectionType == MysqlConnectionConst {
		connection, err := NewMysqlDatabaseConnection()

		if err != nil {
			return nil, err
		}

		return connection, nil
	}

	if databaseFactory.ConnectionType == SqliteConnectionConst {
		connection, err := NewSqliteDatabaseConnection()

		if err != nil {
			return nil, err
		}

		return connection, nil
	}

	return nil, errors.New(InvalidConnectionConst)
}

package database

import (
	"database/sql"
	port "socialMedia/application/port/database"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase(connectionType string) (port.DatabaseConnectionInterface, error) {
	databaseFactory := NewDatabaseFactory(connectionType)

	connection, err := databaseFactory.MakeInstance()

	if err != nil {
		return nil, err
	}

	return connection, nil
}

package database

import (
	"database/sql"
	port "socialMedia/application/port/database"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteDriverName     = "sqlite3"
	sqliteDataSourceName = ":memory:"
)

type SqliteDatabaseConnection struct {
	connection *sql.DB
}

func NewSqliteDatabaseConnection() (port.DatabaseConnectionInterface, error) {
	return &SqliteDatabaseConnection{}, nil
}

func (sqliteDatabaseConnection *SqliteDatabaseConnection) OpenConnection() error {
	sqlite, err := sql.Open(sqliteDriverName, sqliteDataSourceName)

	if err != nil {
		return err
	}

	sqliteDatabaseConnection.connection = sqlite

	return nil
}

func (sqliteDatabaseConnection *SqliteDatabaseConnection) GetConnection() *sql.DB {
	if sqliteDatabaseConnection.connection == nil {
		if err := sqliteDatabaseConnection.OpenConnection(); err != nil {
			panic(err)
		}
	}

	return sqliteDatabaseConnection.connection
}

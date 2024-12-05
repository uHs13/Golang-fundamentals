package userRepository_test

import (
	"socialMedia/infra/database"
	"testing"
)

func TestShouldProperlyCreateAnUser(t *testing.T) {
	factory := database.NewDatabaseFactory(database.SqliteConnectionConst)

	sqliteDb, err := factory.MakeInstance()

	if err != nil {
		panic(err)
	}

	query := "CREATE TABLE user(id string, name string, nickname string, email string, password string)"

	statement, err := sqliteDb.GetConnection().Prepare(query)

	if err != nil {
		panic(err)
	}

	if _, err := statement.Exec(); err != nil {
		panic(err)
	}
}

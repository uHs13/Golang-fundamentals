package userRepository_test

import (
	port "socialMedia/application/port/database"
	"socialMedia/infra/database"
	userRepository "socialMedia/infra/database/repository/user"
	factory "socialMedia/infra/factory/user"
	"testing"

	"github.com/stretchr/testify/require"
)

var SqliteDb port.DatabaseConnectionInterface

func SetUpSqlite() {
	CreateSqliteConnection()
	CreateUserTable()
}

func CreateSqliteConnection() {
	factory := database.NewDatabaseFactory(database.SqliteConnectionConst)

	sqlite, err := factory.MakeInstance()

	SqliteDb = sqlite

	if err != nil {
		panic(err)
	}
}

func CreateUserTable() {
	query := "CREATE TABLE user(id string, name string, nickname string, email string, password string)"

	statement, err := SqliteDb.GetConnection().Prepare(query)

	if err != nil {
		panic(err)
	}

	if _, err := statement.Exec(); err != nil {
		panic(err)
	}
}

func TestShouldProperlyCreateAnUser(t *testing.T) {
	SetUpSqlite()

	createUserRepository := userRepository.NewCreateUserRepository(SqliteDb)

	user, err := factory.NewUserFactory().MakeComplete(
		"John Cena",
		"john.cena",
		"john.cena@email.com",
		"A23#asdD",
	)

	require.Nil(t, err)

	createdUser, err := createUserRepository.Execute(user)

	require.Nil(t, err)
	require.Equal(t, user.GetId(), createdUser.GetId())
	require.Equal(t, user.GetName(), createdUser.GetName())
	require.Equal(t, user.GetNickname(), createdUser.GetNickname())
	require.Equal(t, user.GetEmail(), createdUser.GetEmail())
}

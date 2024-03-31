package main

import (
	"database/src/adapter/http"
	database "database/src/infra/database"
	"database/src/infra/dotenv"
)

func main() {
	if err := dotenv.Load(); err != nil {
		panic(err)
	}

	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := http.NewHttpServer(db)

	if err := server.Start(); err != nil {
		panic(err)
	}
}

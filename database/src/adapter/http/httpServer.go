package http

import (
	"database/sql"
	user "database/src/adapter/http/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	httpServerPort = "HTTP_SERVER_PORT"
)

type HttpServer struct {
	app        *mux.Router
	connection *sql.DB
}

func NewHttpServer(connection *sql.DB) *HttpServer {
	return &HttpServer{
		mux.NewRouter(),
		connection,
	}
}

func (httpServer *HttpServer) Start() error {
	if err := httpServer.registerRoutes(); err != nil {
		return err
	}

	if err := httpServer.initialize(); err != nil {
		return err
	}

	return nil
}

func (httpServer *HttpServer) registerRoutes() error {
	user.NewUserRoutes(
		httpServer.app,
		httpServer.connection,
	).Register()

	return nil
}

func (httpServer *HttpServer) initialize() error {
	portString := fmt.Sprintf(":%s", os.Getenv(httpServerPort))
	fmt.Println("Http server is ready! " + portString)
	err := http.ListenAndServe(portString, httpServer.app)

	if err != nil {
		return err
	}

	return nil
}

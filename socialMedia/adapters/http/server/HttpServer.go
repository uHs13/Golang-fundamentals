package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (httpServer *HttpServer) Start() error {
	muxRouter := mux.NewRouter()
	// negroniMiddleware := negroni.New(negroni.NewLogger())

	// handler.MakeProductHandlers(muxRouter, negroniMiddleware, httpServer.Service)
	http.Handle("/", muxRouter)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":3000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	fmt.Println(`The http server has been started - port ":3000"`)

	err := server.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"./driver"
	handlerDomain "./handler/http"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
)

const (
	port = ":3000"
	dbName = "testzofi_db"
	dbHost = "localhost"
	dbPort = "26257"
)

func main() {

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	handler := handlerDomain.NewDomainHandler(connection)

	router.Route("/", func(rt chi.Router) {
		rt.Mount("/domains", initRouter(handler))
	})
	http.ListenAndServe(port, router)
}

func initRouter(handler *handlerDomain.Domain) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handler.GetS)
	router.Get("/host={host}", handler.GetByHost)

	return router
}


package main

import (
	"errors"
	"log"
	"net/http"

	"pokeapi/api/internal/domain"
	"pokeapi/api/internal/repository/client"
	"pokeapi/api/internal/repository/storage"
	"pokeapi/api/internal/rest"
)

const defaultPort = "8080" // TODO move to config

func main() {
	// init dependencies
	st := storage.NewRepository()
	cl := client.NewRepository()
	service := domain.NewService(st, cl)

	// start server
	log.Printf("starting server on :%s", defaultPort)
	if err := http.ListenAndServe(":"+defaultPort, rest.NewServer(service)); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error starting server: %s\n", err)
			return
		}
		log.Println("server stopped")
	}
}

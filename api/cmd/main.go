package main

import (
	"errors"
	"log"
	"net/http"
	"pokeapi/api/cmd/args"
	"pokeapi/api/internal/config"

	"pokeapi/api/internal/domain"
	"pokeapi/api/internal/repository/client"
	"pokeapi/api/internal/repository/storage"
	"pokeapi/api/internal/rest"
)

func main() {
	// init config variables
	args.ParseArgs()
	conf := config.Get()

	// init dependencies
	st := storage.NewRepository(conf.Storage)
	cl := client.NewRepository(conf.RestClient)
	service := domain.NewService(st, cl)

	// start server
	log.Printf("starting server on :%s", args.Port)
	if err := http.ListenAndServe(":"+args.Port, rest.NewServer(service)); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error starting server: %s\n", err)
			return
		}
		log.Println("server stopped")
	}
}

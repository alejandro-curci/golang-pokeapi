package rest

import (
	"net/http"

	"pokeapi/api/internal/domain/entities"
)

type (
	Server struct {
		router  *http.ServeMux
		service Service
	}
	Service interface {
		FetchData(id int) error
		GetPokemon(id int) (entities.Pokemon, error)
	}
)

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func NewServer(srv Service) *Server {
	server := &Server{
		router:  http.NewServeMux(),
		service: srv,
	}

	server.router.HandleFunc("/pokeapi/get", server.get)
	server.router.HandleFunc("/pokeapi/fetch", server.fetch)

	return server
}

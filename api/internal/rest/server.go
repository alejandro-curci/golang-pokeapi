package rest

import (
	"errors"
	"net/http"

	"pokeapi/api/internal/domain/entities"
	apiErr "pokeapi/api/internal/errors"
)

type (
	Server struct {
		router  *http.ServeMux
		service Service
	}
	Service interface {
		FetchData(id int) error
		GetSummary(id int) (entities.Summary, error)
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

func Error(w http.ResponseWriter, err error) {
	var e apiErr.ApiError
	if ok := errors.As(err, &e); !ok {
		e = apiErr.ErrUnhandled
	}
	http.Error(w, e.Error(), e.Status())
}

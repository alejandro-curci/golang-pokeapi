package rest

import (
	"encoding/json"
	"net/http"

	"pokeapi/api/internal/errors"
)

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Error(w, errors.ErrBadRequest)
		return
	}
	id, err := IDFromParams(r)
	if err != nil {
		Error(w, err)
		return
	}
	pokemon, err := s.service.GetPokemon(id)
	if err != nil {
		Error(w, err)
		return
	}
	pokeBytes, err := json.Marshal(&pokemon)
	if err != nil {
		Error(w, errors.ErrMarshal)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pokeBytes)
}

func (s *Server) fetch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Error(w, errors.ErrBadRequest)
		return
	}
	id, err := IDFromParams(r)
	if err != nil {
		Error(w, err)
		return
	}
	if err := s.service.FetchData(id); err != nil {
		Error(w, err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pokeapi/api/internal/errors"
)

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Error(w, errors.ErrBadRequest)
	}
	idStr := r.URL.Query()["id"][0]
	if idStr == "" {
		Error(w, errors.ErrBadRequest)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, errors.ErrBadRequest)
	}
	pokemon, err := s.service.GetSummary(int(id))
	if err != nil {
		Error(w, err)
	}
	pokeBytes, err := json.Marshal(&pokemon)
	if err != nil {
		Error(w, errors.ErrMarshal)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pokeBytes)
}

func (s *Server) fetch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Error(w, errors.ErrBadRequest)
	}
	idStr := r.URL.Query()["id"][0]
	if idStr == "" {
		Error(w, errors.ErrBadRequest)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, errors.ErrBadRequest)
	}
	if err := s.service.FetchData(int(id)); err != nil {
		Error(w, err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

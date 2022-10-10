package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"pokeapi/api/internal/domain/entities"
)

const requestURL = "https://pokeapi.co/api/v2/pokemon/%d"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) Find(id int) (entities.Pokemon, error) {
	url := fmt.Sprintf(requestURL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return entities.Pokemon{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entities.Pokemon{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.Pokemon{}, err
	}
	var pokemon entities.Pokemon
	return pokemon, json.Unmarshal(body, &pokemon)
}

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokeapi/api/internal/config"

	"pokeapi/api/internal/domain/entities"
)

const requestURL = "/pokemon/%d"

type Repository struct {
	endpoint string
}

func NewRepository(conf config.RestClient) *Repository {
	return &Repository{
		endpoint: conf.URL + requestURL,
	}
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

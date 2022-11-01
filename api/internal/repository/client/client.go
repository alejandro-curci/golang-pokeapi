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

func (r Repository) Find(id int) (entities.PokeData, error) {
	url := fmt.Sprintf(r.endpoint, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return entities.PokeData{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entities.PokeData{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.PokeData{}, err
	}
	var pokemon entities.PokeData
	return pokemon, json.Unmarshal(body, &pokemon)
}

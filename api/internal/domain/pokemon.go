package domain

import (
	"pokeapi/api/internal/domain/entities"
	"pokeapi/api/internal/errors"
)

type (
	Service struct {
		storage Storage
		client  Client
	}
	Storage interface {
		Get(id int) (entities.Pokemon, error)
		Save(pokemon entities.Pokemon) error
	}
	Client interface {
		Find(id int) (entities.PokeData, error)
	}
)

func NewService(storage Storage, client Client) *Service {
	return &Service{storage: storage, client: client}
}

// FetchData makes a request to the data source and saves the response into the storage
func (s *Service) FetchData(id int) error {
	p, err := s.client.Find(id)
	if err != nil {
		return errors.ErrRestClient
	}
	if p.IsEmpty() {
		return errors.ErrNotFound
	}
	if err := s.storage.Save(p.ToPokemon()); err != nil {
		return errors.ErrStorage
	}
	return nil
}

// GetPokemon returns a summary from the stored data
func (s *Service) GetPokemon(id int) (entities.Pokemon, error) {
	p, err := s.storage.Get(id)
	if err != nil {
		return entities.Pokemon{}, errors.ErrStorage
	}
	if (p == entities.Pokemon{}) {
		return entities.Pokemon{}, errors.ErrNotFound
	}
	return p, nil
}

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
		Save(id int, pokemon entities.Pokemon) error
	}
	Client interface {
		Find(id int) (entities.Pokemon, error)
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
	if err := s.storage.Save(id, p); err != nil {
		return errors.ErrStorage
	}
	return nil
}

// GetSummary returns a summary from the stored data
func (s *Service) GetSummary(id int) (entities.Summary, error) {
	p, err := s.storage.Get(id)
	if err != nil {
		return entities.Summary{}, errors.ErrStorage
	}
	if p.IsEmpty() {
		return entities.Summary{}, errors.ErrNotFound
	}
	return p.ToSummary(), nil
}

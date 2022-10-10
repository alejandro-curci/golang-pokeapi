package storage

import "pokeapi/api/internal/domain/entities"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) Get(id int) (entities.Pokemon, error) {
	return entities.Pokemon{}, nil
}

func (r Repository) Save(id int, pokemon entities.Pokemon) error {
	return nil
}

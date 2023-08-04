package service

import (
	"to-do-movie_list/internal/database"
	"to-do-movie_list/internal/domain"
)

type ServiceCrudOperations struct {
	CrudOperations
}

func NewServiceCrudOperations(database *database.Database) *ServiceCrudOperations {
	return &ServiceCrudOperations{CrudOperations: database.CrudOperations}
}

func (s *ServiceCrudOperations) GetAll() ([]domain.Movie, error) {
	return s.CrudOperations.GetAll()
}

func (s *ServiceCrudOperations) GetById(id int) domain.Movie {
	return s.CrudOperations.GetById(id)
}

func (s *ServiceCrudOperations) Post(movie domain.Movie) error {
	return s.CrudOperations.Post(movie)
}

func (s *ServiceCrudOperations) UpdateById(id int, movie domain.Movie) error {
	return s.CrudOperations.UpdateById(id, movie)
}

func (s *ServiceCrudOperations) DeleteById(id int) error {
	return s.CrudOperations.DeleteById(id)
}

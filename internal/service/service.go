package service

import (
	"to-do-movie_list/internal/database"
	"to-do-movie_list/internal/domain"
)

type Service struct {
	CrudOperations
}

type CrudOperations interface {
	GetAll() ([]domain.Movie, error)
	GetById(id int) domain.Movie
	Post(movie domain.Movie) error
	UpdateById(id int, movie domain.Movie) error
	DeleteById(id int) error
}

func NewService(_database *database.Database) *Service {
	return &Service{
		CrudOperations: NewServiceCrudOperations(_database),
	}
}

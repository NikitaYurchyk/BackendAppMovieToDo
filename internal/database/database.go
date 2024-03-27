package database

import (
	"database/sql"
	"to-do-movie_list/internal/domain"
)

type Database struct {
	CrudOperations
}

type CrudOperations interface {
	GetAll() ([]domain.Movie, error)
	GetById(id int) domain.Movie
	Post(movie domain.Movie) error
	UpdateById(id int, movie domain.Movie) error
	DeleteById(id int) error
}

func NewDB(db *sql.DB) *Database {
	return &Database{CrudOperations: NewDbCrudOperations(db)}
}

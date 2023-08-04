package database

import (
	"database/sql"
	"to-do-movie_list/internal/domain"
)

type Database struct {
	CrudOperations
	UsersOperations
}

type UsersOperations interface {
	Create(user domain.User) error
	GetByCred(email string, password string) (domain.User, error)
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

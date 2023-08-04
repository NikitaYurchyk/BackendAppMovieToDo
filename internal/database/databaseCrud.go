package database

import (
	"database/sql"
	"fmt"
	"to-do-movie_list/internal/domain"
)

type DbCrudOperations struct {
	db *sql.DB
}

func NewDbCrudOperations(_database *sql.DB) *DbCrudOperations {
	return &DbCrudOperations{db: _database}

}

func (d *DbCrudOperations) GetAll() ([]domain.Movie, error) {
	var movies []domain.Movie
	rows, err := d.db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var movie domain.Movie
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Director, &movie.Year, &movie.Country); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, err
}

func (d *DbCrudOperations) GetById(id int) domain.Movie {
	var movie domain.Movie
	row := d.db.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(&movie.Id, &movie.Title, &movie.Director, &movie.Year, &movie.Country)
	fmt.Println(row)

	if row != nil {
		return domain.Movie{}
	}

	return movie
}

func (d *DbCrudOperations) Post(movie domain.Movie) error {

	tx, err := d.db.Begin()
	if err != nil {
		defer tx.Rollback()
		return nil
	}
	_, err = tx.Exec("insert into movies (title, director, year, country) values ($1, $2, $3, $4)", movie.Title, movie.Director, movie.Year, movie.Country)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}
	return nil
}

func (d *DbCrudOperations) UpdateById(id int, movie domain.Movie) error {
	tx, err := d.db.Begin()
	if err != nil {
		defer tx.Rollback()
		return nil
	}
	_, err = tx.Exec("update movies set title=$1, director=$2, year=$3, country=$4 where id = $5", movie.Title, movie.Director, movie.Year, movie.Country, id)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}

	return nil
}

func (d *DbCrudOperations) DeleteById(id int) error {
	tx, err := d.db.Begin()
	if err != nil {
		defer tx.Rollback()
		return nil
	}
	_, err = tx.Exec("delete from movies where id = $1", id)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}

	return nil
}

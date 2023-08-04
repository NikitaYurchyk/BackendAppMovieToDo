package database

import (
	"database/sql"
	"errors"
	"to-do-movie_list/internal/domain"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (u *Users) Create(user domain.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		defer tx.Rollback()
		return err
	}

	_, err = tx.Exec("insert into users (name, email, password, registered_at) values ($1, $2, $3, $4)", user.Name, user.Email, user.Password, user.RegistAt)
	if err != nil {
		defer tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}

	return nil
}

func (u *Users) GetByCred(email string, password string) (domain.User, error) {

	var user domain.User
	row := u.db.QueryRow("select * from users where email = $1 and password = $2").Scan(user.Email, user.Password)

	if row != nil {
		return user, errors.New(row.Error())
	}
	return user, nil
}

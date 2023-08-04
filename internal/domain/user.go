package domain

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	RegistAt time.Time `json:"registered_at"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s SignInInput) Validate() error {
	return validate.Struct(s)
}

func (s SignUpInput) Validate() error {
	return validate.Struct(s)
}

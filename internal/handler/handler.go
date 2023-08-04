package handler

import (
	"github.com/go-chi/chi"
	"to-do-movie_list/internal/domain"
	"to-do-movie_list/internal/service"
)

type CrudOperations interface {
	GetAll() ([]domain.Movie, error)
	GetById(id int) domain.Movie
	Post(movie domain.Movie) error
	UpdateById(id int, movie domain.Movie) error
	DeleteById(id int) error
}

type Handler struct {
	CrudOperations
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{CrudOperations: NewHandlerCrudOperations(service)}
}

func (h *Handler) InitHandler() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(logMethod)
	mux.Route("/api", func(r chi.Router) {
		r.Get("/", h.GetAllMovies)
		r.Post("/", h.PostMovie)
		r.Route("/{movieId}", func(r chi.Router) {
			r.Get("/", h.GetMovieById)
			r.Put("/", h.UpdateMovieById)
			r.Delete("/", h.DeleteMovieById)
		})
	})
	return mux
}

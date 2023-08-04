package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"strconv"
	"to-do-movie_list/internal/domain"
	"to-do-movie_list/internal/service"
)

type HandlerCrudOperations struct {
	CrudOperations
}

func NewHandlerCrudOperations(service *service.Service) *HandlerCrudOperations {
	return &HandlerCrudOperations{
		CrudOperations: service.CrudOperations,
	}

}

func (h *Handler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/json")

	resp, err := json.Marshal(movies)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *Handler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "movieId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		fmt.Println("HERE1")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	movie := h.GetById(id)

	w.Header().Set("Content-Type", "text/json")

	resp, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("HERE3")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *Handler) PostMovie(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var movie domain.Movie

	if err := json.Unmarshal(data, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = h.Post(movie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handler) UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "movieId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	movie := h.GetById(id)

	updMovie, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(updMovie, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.UpdateById(id, movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "movieId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = h.GetById(id)

	err = h.DeleteById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

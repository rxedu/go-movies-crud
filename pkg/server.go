package pkg

import (
	"github.com/gorilla/mux"
	"github.com/rxedu/go-movies-crud/v1/internal"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", internal.GetMovies).Methods("GET")
	r.HandleFunc("/movies", internal.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", internal.GetMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", internal.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", internal.DeleteMovie).Methods("DELETE")
	return r
}

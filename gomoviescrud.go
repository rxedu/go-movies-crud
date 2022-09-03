package gomoviescrud

import (
	"github.com/gorilla/mux"
	"github.com/rxedu/go-movies-crud/internal/crud"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", crud.GetMovies).Methods("GET")
	r.HandleFunc("/movies", crud.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", crud.GetMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", crud.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", crud.DeleteMovie).Methods("DELETE")
	crud.InitData()
	return r
}

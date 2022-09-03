package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rxedu/go-movies-crud"
)

func main() {
	s := gomoviescrud.NewServer()
	fmt.Printf("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", s))
}

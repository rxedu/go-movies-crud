package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rxedu/go-movies-crud/v1/pkg"
)

func main() {
	s := pkg.NewServer()
	fmt.Printf("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", s))
}

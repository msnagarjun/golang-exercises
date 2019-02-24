package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers)
	r.HandleFunc("/users/{id}", GetUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}

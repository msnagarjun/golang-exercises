package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers : Get users from database
func GetUsers(w http.ResponseWriter, request *http.Request) {
	users := getUsers()
	json.NewEncoder(w).Encode(users)
}

// GetUser : Get user from database
func GetUser(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		panic(err)
	}
	users := getUser(int32(id))
	json.NewEncoder(w).Encode(users)
}

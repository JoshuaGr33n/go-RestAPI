package main

import (
	// "fmt"
	"net/http"
	"RestAPI/endpoints"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", endpoints.GetUsers).Methods("GET")
	http.ListenAndServe(":8080", router)
}


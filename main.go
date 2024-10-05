package main

import (
	"merchantBank/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/login", handlers.Login).Methods("POST")
    r.HandleFunc("/payment", handlers.Payment).Methods("POST")
    r.HandleFunc("/logout", handlers.Logout).Methods("POST")

    http.ListenAndServe(":8080", r)
}

package main

import (
	"crud/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/create", services.CreateUser).Methods("POST")
	router.HandleFunc("/user", services.SearchUsers).Methods("GET")
	router.HandleFunc("/user/{id}", services.SearchUser).Methods("GET")
	router.HandleFunc("/user/{id}", services.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", services.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

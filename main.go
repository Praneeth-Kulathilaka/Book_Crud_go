package main

import (
	"fmt"
	"net/http"
	"BookApi/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", handlers.GetABook).Methods("GET")
	r.HandleFunc("/books/{id}",handlers.Update).Methods("PUT")
	r.HandleFunc("/books/{id}",handlers.Delete).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
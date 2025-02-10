package main

import (
	"BookApi/config"
	"BookApi/handlers"
	"fmt"
	"log"
	"net/http"

	// "BookApi/handlers/external"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	client := config.InitRedis()

	if client == nil {
		log.Println("Error getting redis client ")
		return
	}

	// err := client.Set(config.Ctx, "testKey1", "testValue1", 0).Err()
	// if err != nil {
	// 	log.Println("Error seting key value", err)
	// 	return
	// }



	r.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", handlers.GetABook).Methods("GET")
	r.HandleFunc("/books/{id}",handlers.Update).Methods("PUT")
	r.HandleFunc("/books/{id}",handlers.Delete).Methods("DELETE")
	r.HandleFunc("/onlineBooks",handlers.Fetch).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
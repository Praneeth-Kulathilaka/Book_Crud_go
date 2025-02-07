package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Book struct {
	ID int `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []*Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	log.Println("Recieved", book)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println(err)
		return
	}
	book.ID = uuid.New().ClockSequence()
	
	books = append(books, book)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(book)
}
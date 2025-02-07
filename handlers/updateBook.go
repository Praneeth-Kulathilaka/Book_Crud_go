package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	log.Println("Log id of put ", id)
	var requestedBook Book
	if err != nil || id == 0 {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&requestedBook)
	log.Println("Recieved", requestedBook)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println(err)
		return
	}
	
	for _, book := range books {
		if book.ID == id {
			
			if requestedBook.Author != "" && requestedBook.Title != "" {
				book.Author = requestedBook.Author
				book.Title = requestedBook.Title
			} else if requestedBook.Author != "" {
				book.Author = requestedBook.Author
			} else {
				book.Title = requestedBook.Title
			}
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)

}
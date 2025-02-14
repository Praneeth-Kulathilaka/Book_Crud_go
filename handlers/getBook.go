package handlers

import (
	"BookApi/config"
	"BookApi/handlers/channels"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	client := config.GetRedisClient()

	channels.SendLogMessage("GET","Get books Api called")

	keys, err := client.Keys(config.Ctx, "books:*").Result()
	if err != nil {
		log.Println("Failed to fetch keys",err)
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}
	var books []Book
	for _, key := range keys {
		bookData, err := client.HGetAll(config.Ctx, key).Result()
		if err != nil || len(bookData) == 0 {
			continue
		}
		id, _ := strconv.Atoi(bookData["id"])
		book := Book{
			ID: id,
			Title: bookData["title"],
			Author: bookData["author"],
		}
		books = append(books, book)
	}

	responseJSON, _ := json.Marshal(books)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func GetABook (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	log.Println("Put id is ",id)
	if err != nil || id == 0 {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)

}
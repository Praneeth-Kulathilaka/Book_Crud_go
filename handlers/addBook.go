package handlers

import (
	"BookApi/config"
	"BookApi/handlers/channels"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Book struct {
	ID int `json:"id" redis:"id"`
	Title  string `json:"title" redis:"title"`
	Author string `json:"author" redis:"author"`
}

var books []*Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	// log.Println("Recieved", book)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println(err)
		return
	}
	book.ID = uuid.New().ClockSequence()

	message := fmt.Sprintf("/books -- ID: %d", book.ID)
	channels.SendLogMessage("POST",message)

	client := config.GetRedisClient()

	// data, _ := json.Marshal(book)
	// _, err = client.SAdd(config.Ctx, "books_set", data, 0).Result()
	data := map[string]interface{}{
		"id": book.ID,
		"title": book.Title,
		"author": book.Author,
	} 
	// log.Println("Dataq: ",data)
	hashKeyValue := fmt.Sprintf("books:%d",book.ID)
	_, err = client.HSet(config.Ctx,hashKeyValue,data).Result()
	if err != nil {
		log.Println("Error caching data",err)
		return
	}
	
	books = append(books, book)
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(book)
}
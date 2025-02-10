package handlers

import (
	"BookApi/config"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	client := config.GetRedisClient()

	booksInRedis, _ := client.SMembers(config.Ctx, "books_set").Result()
	
	booksArray, _ := json.Marshal(booksInRedis)
	log.Println(booksArray)
	responseJSON := "[" + strings.Join(booksInRedis, ",") + "]"

	w.Header().Set("Content-Type","application/json")
	w.Write([]byte(responseJSON))
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
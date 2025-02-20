package handlers

import (
	"BookApi/handlers/channels"
	"fmt"
	// "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	// log.Println("Put id is ",id)
	message := fmt.Sprintf("/books -- ID: %d", id)
	channels.SendLogMessage("DELETE",message)

	
	if err != nil || id == 0 {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == id {
			books = append( books[:id], books[id+1:]... )
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)

}
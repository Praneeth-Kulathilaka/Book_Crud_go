package handlers

import (
	"BookApi/handlers/external"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
)
	

func Fetch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("isbn")
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")
	log.Println("Isbn from query: ",id)
	log.Println("Title from query: ",title)

	var response []byte
	var err error

	if title != "" {
		response, err = external.FetchByTitle(title)
	} else if author != ""{

	} else {
		response, err = external.FetchByID(id)
	}
	if err != nil {
		log.Println("Error fetching data %s",err)
	}
	// log.Println("Response to fetchBooks: ",response)

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println("Error fetching data %s",err)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(result)


}
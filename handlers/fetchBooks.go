package handlers

import (
	"BookApi/handlers/channels"
	"BookApi/handlers/dto"
	"BookApi/handlers/external"
	"encoding/json"
	"log"
	"net/http"
)

// type BookDTO struct {
// 	Author []string `json:"author_name"`
// 	Title string `json:"title"`
// }

type ApiResponse struct {
	Docs []dto.BookDTO `json:"docs"`
}
// type ApiResponse2 struct {
// 	ISBN []dto.ByIdDTO `json:"ISBN"`
// }

func Fetch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("isbn")
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")
	// log.Println("Isbn from query: ",id)
	// log.Println("Title from query: ",title)
	// log.Println("Author from query: ",author)

	channels.SendLogMessage("GET","/onlineBooks")


	var response []byte
	var err error

	if title != "" {
		response, err = external.FetchByTitle(title)
	} else if author != ""{
		response, err = external.FetchByAuthor(author)
	} else {
		response, err = external.FetchByID(id)
	}
	if err != nil {
		log.Printf("Error fetching data %s",err)
	}

	// var result2 ApiResponse2
	// if id != ""{
	// 	log.Println("Response",response)
	// 	err = json.Unmarshal(response, &result2)
	// 	if err != nil {
	// 		log.Printf("Error fetching data %s",err)
	// 	}
	// 	log.Println("First book...",result2)

	// 	w.Header().Set("Content-Type","application/json")
	// 	json.NewEncoder(w).Encode(result2)
	// 	return
	// }

	// var result map[string]interface{}
	var result ApiResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Printf("Error fetching data %s",err)
	}

	log.Println("First book...",result)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(result)


}
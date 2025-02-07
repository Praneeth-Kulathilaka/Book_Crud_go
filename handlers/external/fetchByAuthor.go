package external

import (
	// "fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func FetchByAuthor(author string) ([]byte, error) {
	log.Println("Isbn called: ",author)

	titleToSend := strings.ReplaceAll(author," ", "+")
	log.Println("TitletoSend",titleToSend)


	response, err := http.Get("https://openlibrary.org/search.json?title="+titleToSend)
	if err != nil {
		log.Println("Error: ",err)
	}
	log.Println("Response from openlibrary: ",response)
	log.Println("Response body: ",response.Body)
	responseData, _ := io.ReadAll(response.Body)
	
	return responseData, err
}
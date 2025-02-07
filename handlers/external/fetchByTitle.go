package external

import (
	// "fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func FetchByTitle(title string) ([]byte, error) {
	log.Println("Isbn called: ",title)

	titleToSend := strings.ReplaceAll(title," ", "+")
	log.Println("TitletoSend",titleToSend)


	response, err := http.Get("https://openlibrary.org/search.json?title="+titleToSend)
	if err != nil {
		log.Println("Error: ",err)
	}
	log.Println("Response from openlibrary: ",response)
	log.Println("Response body: ",response.Body)
	responseData, _ := io.ReadAll(response.Body)
	log.Println("Response data",responseData)
	return responseData, err
}
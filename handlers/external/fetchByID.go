package external

import (
	// "fmt"
	"io"
	"log"
	"net/http"
)

func FetchByID(isbn string) ([]byte, error) {
	log.Println("Isbn called: ",isbn)
	// url := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:" +isbn+"&format=json&jscmd=data")

	response, err := http.Get("https://openlibrary.org/api/books?bibkeys=ISBN:" +isbn+"&format=json&jscmd=data")
	if err != nil {
		log.Println("Error: ",err)
	}
	log.Println("Response from openlibrary: ",response)
	log.Println("Response body: ",response.Body)
	responseData, _ := io.ReadAll(response.Body)
	
	return responseData, err
}
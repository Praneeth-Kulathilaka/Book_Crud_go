package channels

import (
	"log"
	"sync"
	"time"
)

type LogMessage struct {
	Method   string
	Message string
	Time    time.Time
}

var logChannel = make(chan LogMessage, 100)

func LogListner(wg *sync.WaitGroup) {
	defer wg.Done()
	for logs := range logChannel {
		log.Printf("[%s] %s: %s\n",logs.Time.Format(time.RFC3339), logs.Method, logs.Message)
	}	
}

func SendLogMessage (method, message string){
	logChannel <- LogMessage{
		Method: method,
		Message: message,
		Time: time.Now(),
	}
}
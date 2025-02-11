package redis

import (
	// "BookApi/handlers/redis"
	"log"
	"sync"
	"time"
)

func MonitorIdleTasks(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(10*time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		err := DeleteKeys(100)
		if err != nil {
            log.Println("Error deleting idle keys", err)
        }
		log.Println("Monitor task function called")
	}

}
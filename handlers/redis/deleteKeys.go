package redis

import (
	"BookApi/config"
	// "internal/coverage/rtcov"
	// "errors"
	"log"
	// "time"
)

func DeleteKeys(threshold int64) error {
	var cursor uint64
	client := config.GetRedisClient()

	for {
		keys, nextCursor, err := client.Scan(config.Ctx, cursor, "books:*",100).Result()
		if err != nil {
			log.Println("Error scanning client",err)
			return err
		}
		for _, key := range keys {
			// idleTime, err := client.ObjectIdleTime(config.Ctx, key).Result()
			idleTime, err := client.Do(config.Ctx, "OBJECT", "IDLETIME", key).Int64()
			if err != nil {
				log.Println("Error getting the idle time", err)
				continue
			} 
			// idleTime = time.Duration(idleTime)*time.Second
			if idleTime > threshold {
				_, err := client.Del(config.Ctx, key).Result()
				if err != nil {
					log.Println("Error deleting cache data",err)
					return err
				} else {
					log.Println("Deleted idle key from redif",key)
				}

			}
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}
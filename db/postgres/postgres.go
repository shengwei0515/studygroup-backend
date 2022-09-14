package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

var db *gorm.DB

func InitWithRetry(dbUri, dbDriver string, reconnectTimes int, reconnectBounceSec time.Duration) {
	var err interface{}
	for i := 0; i < reconnectTimes; i += 1 {
		err = initDb(dbUri, dbDriver)
		if err == nil {
			return
		}
		log.Printf("DB Init error %s\nWait for %s second to reconnect DB", err, reconnectBounceSec)
		time.Sleep(reconnectBounceSec)
	}

	log.Panicf("DB Init error after retry %c times, error: %s", reconnectBounceSec, err)
}

func initDb(dbUri, dbDriver string) error {
	postgres, err := gorm.Open(dbDriver, dbUri)
	if err != nil {
		return fmt.Errorf("Connect to postgres failed with error: %s", err)
	}

	db = postgres
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}

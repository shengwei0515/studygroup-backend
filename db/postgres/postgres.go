package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

var db *gorm.DB

const dbReconnectTimes = 5
const dbReconnectSec = 1

func InitWithRetry(dbUri, dbDriver string, reconnectTimes int, reconnectBounceSec time.Duration) {
	var err interface{}
	for i := 0; i < reconnectTimes; i += 1 {
		err = Init(dbUri, dbDriver)
		if err == nil {
			break
		}
		log.Printf("DB Init error %s\nWait for %s second to reconnect DB", err, reconnectBounceSec)
		time.Sleep(reconnectBounceSec)
	}
	if err != nil {
		log.Panicf("DB Init error after retry %c time, error: %s", reconnectBounceSec, err)
	}
}

func Init(dbUri, dbDriver string) error {
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

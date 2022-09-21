package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitWithRetry(dbUri, dbDriver string, reconnectTimes int, reconnectBounceSec time.Duration) {
	var err interface{}
	for i := 0; i < reconnectTimes; i += 1 {
		err = Init(dbUri, dbDriver)
		if err == nil {
			return
		}
		log.Printf("DB Init error %s\nWait for %s second to reconnect DB", err, reconnectBounceSec)
		time.Sleep(reconnectBounceSec)
	}

	log.Panicf("DB Init error after retry %c times, error: %s", reconnectBounceSec, err)
}

func Init(dbUri, dbDriver string) error {
	postgresDb, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Connect to postgres failed with error: %s", err)
	}

	db = postgresDb
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

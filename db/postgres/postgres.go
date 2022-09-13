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

func InitWithRetry() {
	var err interface{}
	for i := 0; i < dbReconnectTimes; i += 1 {
		err = Init()
		if err == nil {
			break
		}
		log.Printf("DB Init error %s", err)
		time.Sleep(time.Duration(int64(dbReconnectSec * time.Second)))
	}
	if err != nil {
		log.Panicf("DB Init error after retry %c time, error: %s", dbReconnectTimes, err)
	}
}

func Init() error {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "password"
	dbName := "postgres"
	dbDriver := "postgres"

	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

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

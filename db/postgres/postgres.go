package postgres

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init() {
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
		log.Panicf("Connect to postgres failed with error %s", err)
		panic("Failed to connect to db")
	}
	db = postgres
}

func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}

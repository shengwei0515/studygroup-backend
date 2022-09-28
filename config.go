package studygroup

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type WebConfig struct {
	DbReconnectTimes     int
	DbReconnectBounceSec time.Duration
	DbUri                string
	DbDriver             string
	ServerAddr           string
	SessionConfig        WebSessionConfig
}

type WebSessionConfig struct {
	RedisConnectSize int
	RedisNetwork     string
	RedisAddr        string
	RedisPassword    string
	SessionName      string
	SessionKey       string
}

func ReadEnvConfig() WebConfig {
	dbReconnectTimes, err := strconv.Atoi(os.Getenv("DB_RECONNECT_TIMES"))
	if err != nil {
		log.Printf("load DB_RECONNECT_TIMES from config fail, use default value 5")
		dbReconnectTimes = 5
	}

	dbReconnectBounceTime, err := strconv.ParseInt(os.Getenv("DB_RECONNECT_BOUNCE_SEC"), 10, 64)
	if err != nil {
		log.Printf("load DB_RECONNECT_BOUNCE_SEC from config fail, use default value 10")
		dbReconnectBounceTime = 10
	}
	dbReconnectBounceSec := time.Duration(dbReconnectBounceTime) * time.Second

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	addr := os.Getenv("SERVER_ADDR")

	sessionKey := os.Getenv("SESSION_KEY")
	redisConnectSize, err := strconv.Atoi(os.Getenv("REDIS_CONNECT_SIZE"))
	if err != nil {
		log.Printf("load REDIS_CONNECT_SIZE from config fail, use default value 10")
		dbReconnectBounceTime = 10
	}

	return WebConfig{
		DbReconnectTimes:     dbReconnectTimes,
		DbReconnectBounceSec: dbReconnectBounceSec,
		DbUri:                dbUri,
		DbDriver:             dbDriver,
		ServerAddr:           addr,
		SessionConfig: WebSessionConfig{
			RedisConnectSize: redisConnectSize,
			RedisNetwork:     os.Getenv("REDIS_NETWORK"),
			RedisAddr:        os.Getenv("REDIS_ADDR"),
			RedisPassword:    os.Getenv("REDIS_PASSWORD"),
			SessionName:      os.Getenv("SESSION_NAME"),
			SessionKey:       sessionKey,
		},
	}
}

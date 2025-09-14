package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Init() *redis.Client {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	password := getEnv("REDIS_PASSWORD", "")
	db, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	log.Printf("Connecting to Redis at %s:%s", host, port)

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error connecting to Redis: %v", err)
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	log.Println("Redis connected successfully")
	return client

}
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetClient() *redis.Client {
	if client == nil {
		client = Init()
	}
	return client
}

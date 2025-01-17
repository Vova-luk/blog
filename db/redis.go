package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// Connecting to Redis
func ConnectToRedis(dbIndex int) (*redis.Client, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_HOST_AND_PORT_REDIS"),
		Password: os.Getenv("DB_PASSWORD_REDIS"),
		DB:       dbIndex,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Redis: %v", err)
	}
	log.Println("Connection to redis")
	return client, nil
}

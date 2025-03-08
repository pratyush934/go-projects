package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNo,
	})
	return client
}

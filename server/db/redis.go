package db

import (
	"os"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

// New : Create a new redis client
func New() *Database {
	return &Database{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})}
}

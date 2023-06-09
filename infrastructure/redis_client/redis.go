package redis_client

import (
	"github.com/redis/go-redis/v9"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

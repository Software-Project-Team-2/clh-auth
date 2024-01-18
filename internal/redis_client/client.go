package redis_client

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func InitClient(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetClient() *redis.Client {
	return client
}

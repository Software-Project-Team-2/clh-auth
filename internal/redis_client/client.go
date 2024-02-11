package redis_client

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitClient() {
	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDb,
	})
}

func GetClient() *redis.Client {
	return client
}

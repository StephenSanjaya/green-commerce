package config

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func SetupRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pw := os.Getenv("REDIS_PW")

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pw,
	})
	return rdb
}

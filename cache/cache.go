package cache

import (
	"gopkg.in/redis.v5"
	"os"
)

var client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URI"),
	Password: "",
	DB:       0,
})

// Set a key in redis
func Set(key, value string) error {
	return client.Set(key, value, 0).Err()
}

// Get a key from redis
func Get(key string) (string, error) {
	return client.Get(key).Result()
}

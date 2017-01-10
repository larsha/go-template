package main

import(
  "gopkg.in/redis.v5"
)

var client = redis.NewClient(&redis.Options{
    Addr: "redis:6379",
    Password: "",
    DB: 0,
});

func Set(key, value string) error {
  return client.Set(key, value, 0).Err()
}

func Get(key string) (string, error) {
  return client.Get(key).Result()
}

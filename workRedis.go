package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

//func getRedis(start string, end string) error {
func getRedis(start string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, start).Result()
	var answer string

	switch err {
	case redis.Nil:
		answer = "key not exist"
	case nil:
		answer = val
	default:
		log.Println("Error take data from Redis", err)
	}
	return answer, nil
}

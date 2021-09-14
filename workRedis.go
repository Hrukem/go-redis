package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"strings"
)

// get function take data from Redis
func get(start int, end int) string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//	res := make([]string, 0)
	res := ""
	for i := start; i <= end; i++ {
		key := "Bora" + strconv.Itoa(i)
		val, err := rdb.Get(context.Background(), key).Result()
		switch err {
		case redis.Nil:
			continue
		case nil:
			//res = append(res, val)
			res = res + val + ", "
		default:
			log.Println("Error get data from Redis", err)
			continue
		}
	}
	//	return res
	return strings.Trim(res, ", ")
}

// put function put data in Redis
func put(str string, t int64) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key := "Bora" + strconv.FormatInt(t, 10)
	fmt.Println("key: ", key)

	err := rdb.Set(context.Background(), key, str, 0).Err()
	return err
}

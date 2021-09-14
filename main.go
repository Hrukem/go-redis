package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func main() {
	if checkConnectRedis() {
		log.Fatal("No connection to Redis")
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/domains/", getData)
	mux.HandleFunc("/links/", inputData)

	fmt.Println("start service on :4010")
	log.Fatal(http.ListenAndServe(":4010", mux))
}

func checkConnectRedis() bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val := rdb.Ping(context.Background())

	if fmt.Sprint(val) != "ping: PONG" {
		return true
	} else {
		return false
	}
}

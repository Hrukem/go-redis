package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// inputData function parses the request, gets a link map of the visited sites,
// processes it, getting a string containing the names of the visited domains and
// puts this string in Redis. The key is the time when the request was received.
func inputData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/links/" && r.Method != "POST" {
		http.NotFound(w, r)
	}

	t := time.Now().Unix()

	var b interface{}
	err1 := json.NewDecoder(r.Body).Decode(&b)
	if err1 != nil || b == nil {
		log.Println("Error NewDecoder in inputData(w,r) line52", err1)
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	m := b.(map[string]interface{})
	i := m["links"]
	s := fmt.Sprint(i)
	s = parseLinks(s)

	err3 := inputRedis(s, t)
	if err3 != nil {
		log.Println("Error input data to Redis", err3)
		_, err4 := fmt.Fprintf(w, "Error get data from Redis:\n %v", err3)
		if err4 != nil {
			log.Println("Error change in inputData(w, r) line.59", err4)
		}
	}

	log.Println("Data insert to Redis")
	_, err5 := fmt.Fprintf(w, "Data saved in Redis, %s", "Ok")
	if err5 != nil {
		log.Println("Error change w in inputData(w, r) line.67", err5)
	}
}

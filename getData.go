package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// getData function parses the request,
// finds the start and end values of the key range,
// receives data from Redis for these keys
// and sends this data in response to the request.
func getData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/domains/" && r.Method != "GET" {
		http.NotFound(w, r)
	}

	// checking that the request body is not empty and the keys contain only numbers
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	start, err11 := strconv.Atoi(from)
	end, err12 := strconv.Atoi(to)
	if err11 != nil || err12 != nil || start > end {
		log.Println("Bad request", err11, err12)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	answer := getRedis(start, end)

	log.Println("Send data from Redis")
	_, err2 := fmt.Fprintf(w, "data from Redis:\n %v", answer)
	if err2 != nil {
		log.Println("Error change w in getData(w, r) line.36", err2)
	}
}

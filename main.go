package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/redis/", workRedis)

	fmt.Println("start service")
	log.Fatal(http.ListenAndServe(":4010", mux))
}

func workRedis(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/redis/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		input(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	//	end := r.URL.Query().Get("end")

	//	err := getRedis(start, end)
	answer, err := getRedis(start)
	if err != nil {
		log.Println("Error get data from Redis", err)
		http.Redirect(w, r, "", 404)
	} else {
		fmt.Println("answer", answer)
		log.Println("Send data from Redis")
		http.Redirect(w, r, "", 200)
	}

}

func input(w http.ResponseWriter, r *http.Request) {
	var b interface{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Body: ", b)
}

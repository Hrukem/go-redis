package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/domains/", getData)
	mux.HandleFunc("/links/", inputData)

	fmt.Println("start service on :4010")
	log.Fatal(http.ListenAndServe(":4010", mux))
}

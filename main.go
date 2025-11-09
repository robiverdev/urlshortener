package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /shorten", shortenHandler)
	mux.HandleFunc("GET /{shortCode}", redirectHandler)

	fmt.Print("Starting server on port :4000\n")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

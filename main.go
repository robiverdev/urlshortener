package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var urlStore = make(map[string]string)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /shorten", shortenHandler)
	mux.HandleFunc("GET /{shortCode}", redirectHandler)

	fmt.Print("Starting server on port :4000\n")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST shortening")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET redirect")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	_, err := io.WriteString(w, "Homepage")

	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

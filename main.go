package main

import (
	"encoding/json"
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

type ShortenRequest struct {
	URL string `json:"url"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	shortCode := "abc123" // TODO: make this random
	urlStore[shortCode] = req.URL
	if err := json.NewEncoder(w).Encode(map[string]string{"short_code": shortCode}); err != nil {
		log.Printf("Error encoding response: %v", err)
	}

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

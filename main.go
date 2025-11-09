package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
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

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // All possible characters for the code
	const length = 6                                                                 // Empty byte array of length 6

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))] // Random number between 0 and charset length
	}
	return string(result) // Convert byte string and return
}

type ShortenRequest struct {
	URL string `json:"url"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	shortCode := generateShortCode()
	urlStore[shortCode] = req.URL
	if err := json.NewEncoder(w).Encode(map[string]string{"short_code": shortCode}); err != nil {
		log.Printf("Error encoding response: %v", err)
	}

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	longURL, exists := urlStore[shortCode]
	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	_, err := io.WriteString(w, "Homepage")

	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

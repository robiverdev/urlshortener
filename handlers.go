package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

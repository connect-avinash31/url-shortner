package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var shortnerService ShortnerService

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Url string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortenedValue, err := shortnerService.ShortenValue(request.Url)
	if err != nil {
		http.Error(w, "Error shortening URL: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"shortened_url": shortenedValue}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func originalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		ShortenedUrl string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	originalValue, err := shortnerService.OriginalValue(request.ShortenedUrl)
	if err != nil {
		http.Error(w, "Error getting original URL: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"original_url": originalValue}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	metrics, err := shortnerService.Metrics()
	if err != nil {
		http.Error(w, "Error getting metrics: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}
func main() {
	// first i will run a http server on the port 8080
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/original", originalHandler)
	http.HandleFunc("/metrics", metricsHandler)
	// first will be writing an url shortner and then running an Server to handle it
	shortnerService := NewUrlShortner()
	// now we can call ShortenValue and OriginalValue methods on urlShortner
	// exmaple
	shortenedValue, err := shortnerService.ShortenValue("www.udemy.com/courses/ai-course")
	if err != nil {
		println("Error shortening value:", err.Error())
		return
	}
	fmt.Println("Shortened Value:", shortenedValue)
	originalValue, err := shortnerService.OriginalValue(shortenedValue)
	if err != nil {
		println("Error getting original value:", err.Error())
		return
	}
	fmt.Println("Original Value:", originalValue)
	// starting http server opn 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println("Error starting server:", err.Error())
		return
	}
}

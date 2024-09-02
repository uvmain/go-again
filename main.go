package main

import (
	"log"
	"net/http"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	redirectMap := map[string]string{
		"/google":     "https://www.google.com",
		"/manchester": "https://my.manchester.ac.uk",
	}

	if redirectUrl, found := redirectMap[r.URL.Path]; found {
		log.Printf("Redirect to %s...\n", redirectMap[r.URL.Path])
		http.Redirect(w, r, redirectUrl, http.StatusFound)
	} else {
		log.Printf("No redirect found for %s...\n", r.URL.Path)
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", redirectHandler)
	port := ":8087"
	log.Printf("Starting server on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "pong"}`)
	})

	fmt.Println("Gin-like server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
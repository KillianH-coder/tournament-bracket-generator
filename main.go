package main

import (
	"tournament/handlers"

	"fmt"
	"net/http"
)

func main() {
	// Handle routes
	http.HandleFunc("/", handlers.IndexHandler)

	// Handle assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	fmt.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"tournament/handlers"

	"fmt"
	"net/http"
)

func main() {
	// Handle routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/simulation", handlers.SimulationHandler)

	// Handle assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	fmt.Println(
		"  __________________________________\n",
		"| Running on http://localhost:8080 |\n",
		"|__________________________________|",
	)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Render a home page. Receive request and build the response.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func main() {
	// Create a handler to facilitate requests to specific parts of a web app.
	http.HandleFunc("/", homePage)

	log.Println("Initializing a web server on port 8080")
	http.ListenAndServe(":8080", nil)

}

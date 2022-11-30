package main

import (
	"fmt"
	"net/http"
)

// Create a basic web server that listens on port 4000.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":4000", nil)

}

package main

import (
	"fmt"
	"net/http"
)

// simple web server that listens on port 4000
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Basic web server.")
	})
	http.ListenAndServe(":4000", nil)
}

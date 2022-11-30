package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// A simple rest api server listening on port 4001 and using chi new router.
func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Server!")
	})
	http.ListenAndServe(":4001", r)
}

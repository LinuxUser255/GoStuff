package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// Create a basic web api server listen on port 4001.
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":4001", r)
}

package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", apiHeader)
	http.ListenAndServe(":8000", r)
}

func apiHeader(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

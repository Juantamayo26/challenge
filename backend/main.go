package main

import (
	"challenge/db"
	"challenge/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	db.Schema()
	r := chi.NewRouter()

	//middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "https://localhost:8081"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)

	//routes
	r.Post("/buyers", routes.CreateBuyer)
	r.Post("/products", routes.CreateProducts)
	r.Post("/transactions", routes.CreateTransactions)

	//r.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8003", r))
}

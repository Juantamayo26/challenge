package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

//Types
type Product struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

type AllProducs []Product

var Products = AllProducs{}

func createProducts(w http.ResponseWriter, r *http.Request) {
	reader := csv.NewReader(r.Body)

	reader.Comma = '\''

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	for _, record := range records {
		product := Product{
			ID:    record[0],
			Name:  record[1],
			Price: record[2],
		}
		Products = append(Products, product)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Products)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "ID")

	for _, p := range Products {
		if p.ID == productID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	fmt.Fprintf(w, "Invalid ID")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", getProducts)
	r.Post("/products", createProducts)
	r.Get("/products/{ID}", getOneProduct)

	log.Fatal(http.ListenAndServe(":8000", r))

}

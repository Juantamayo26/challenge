package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
)

//Types
type Product struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

type AllProducs []Product

var Products = AllProducs{}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\''

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
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
	records, err := readData("products.csv")

	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		product := Product{
			ID:    record[0],
			Name:  record[1],
			Price: record[2],
		}
		Products = append(Products, product)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", getProducts)
	r.Get("/products/{ID}", getOneProduct)

	log.Fatal(http.ListenAndServe(":8000", r))

}

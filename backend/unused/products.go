package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

//Types
type Product struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

type AllProducs []Product

var Products = AllProducs{}

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

//func addToDb(){
//}

func createProducts(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("data")
	fileName := handler.Filename

	outfile, err := os.Create("./temp/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outfile.Close()
	cpy, err := io.Copy(outfile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(cpy)

	csvFile, err := os.Open("./temp/" + fileName)
	if err != nil {
		fmt.Fprintf(w, "Error opening csvFile")
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
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
	////
	///////////////v := reflect.ValueOf(Products[0])
	///////////////typeOfS := v.Type()
	///////////////fmt.Println(typeOfS)

	///////////////for i := 0; i < v.NumField(); i++ {
	///////////////	fmt.Printf("Field: %s\tValue: %q\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	///////////////}
	////
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Products)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", getProducts)
	r.Post("/products", createProducts)
	r.Get("/products/{ID}", getOneProduct)

	log.Fatal(http.ListenAndServe(":8001", r))

}

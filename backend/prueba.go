package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

// Types
type buyer struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type allBuyers []buyer

var buyers = allBuyers{
	{
		ID:   "12321",
		Name: "Task one",
		Age:  12,
	},
}

func getBuyers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers)
}

func getOneBuyer(w http.ResponseWriter, r *http.Request) {
	buyerID := chi.URLParam(r, "ID")

	fmt.Println(buyerID)
	for _, b := range buyers {
		if b.ID == buyerID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	fmt.Fprintf(w, "Invalid ID")
}

func createBuyer(w http.ResponseWriter, r *http.Request) {
	var newBuyer buyer
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal(reqBody, &newBuyer)

	buyers = append(buyers, newBuyer)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBuyer)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/buyers", getBuyers)
	r.Post("/buyers", createBuyer)
	r.Get("/buyers/{ID}", getOneBuyer)
	r.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	//r.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(":8000", r))
}

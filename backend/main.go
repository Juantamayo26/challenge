package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
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

var buyers = allBuyers{}

func getBuyers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers)
}

func getOneBuyer(w http.ResponseWriter, r *http.Request) {
	buyerID := chi.URLParam(r, "ID")
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
	mr, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc := []buyer{}
	for {
		part, err := mr.NextPart()

		// This is OK, no more parts
		if err == io.EOF {
			break
		}

		// Some error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if part.FormName() == "data" {
			jsonDecoder := json.NewDecoder(part)
			err = jsonDecoder.Decode(&doc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	var newBuyer []buyer

	data, err := json.Marshal(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.Unmarshal(data, &newBuyer)

	//buyers = doc //RARO
	buyers = newBuyer

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
	log.Fatal(http.ListenAndServe(":8001", r))
}

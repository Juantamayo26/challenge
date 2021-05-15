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

//types
type Transaction struct {
	ID         string `json:"ID"`
	BuyerID    string `json:"BuyerID"`
	IP         string `json:"IP"`
	Device     string `json:"Device"`
	ProductsID string `json:"ProductsID"`
}

type allTransactions []Transaction

var a = [5]string{"2", "4", "6", "8", "10"}

var Transactions = allTransactions{}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Transactions)
}

func getOneTransaction(w http.ResponseWriter, r *http.Request) {
	TransactionID := chi.URLParam(r, "ID")
	for _, t := range Transactions {
		if t.ID == TransactionID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	fmt.Fprintf(w, "Invalid ID")
}

func main() {
	content, err := ioutil.ReadFile("transactions")
	if err != nil {
		log.Fatal(err)
	}

	var aux string
	var record [5]string
	var index int

	for i := 0; i < len(content); i++ {
		if content[i] != 0 {
			aux = aux + string(content[i])
		}
		if content[i] == 0 && content[i+1] == 0 {

			transaction := Transaction{
				ID:         record[0],
				BuyerID:    record[1],
				IP:         record[2],
				Device:     record[3],
				ProductsID: aux,
			}
			Transactions = append(Transactions, transaction)

			index = 0
			aux = ""
		} else if content[i] == 0 {
			record[index] = aux

			index++
			aux = ""
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/transactions", getTransactions)
	r.Get("/transactions/{ID}", getOneTransaction)
	log.Fatal(http.ListenAndServe(":8000", r))
}

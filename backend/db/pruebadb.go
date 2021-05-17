package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"log"
	"net/http"

	"context"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// If omitempty is not set, then edges with empty values (0 for int/float, "" for string, false
// for bool) would be created for values not specified explicitly.

type Person struct {
	Uid   string  `json:"uid,omitempty"`
	Buyer []buyer `json:"data,omitempty"`
}

// Types
type buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type allBuyers []buyer

var buyers = allBuyers{}

func createBuyer(w http.ResponseWriter, r *http.Request) {
	var newBuyer []buyer

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &newBuyer)
	buyers = newBuyer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBuyer)

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	p := Person{
		Uid:   "_:alice",
		Buyer: newBuyer,
	}
	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		age: int .
		married: bool .
		loc: geo .
	`

	ctx := context.Background()
	err = dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(p)
	log.Println(string(pb))
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	// Assigned uids for nodes which were created would be returned in the assigned.Uids map.
	variables := map[string]string{"$id1": assigned.Uids["alice"]}
	//	q := `query Me($id1: string){
	//		me(func: uid($id1)) {
	//			name
	//			age
	//			data @filter(eq(name, "Apicella")){
	//				id
	//				name
	//				age
	//			}
	//		}
	//	}`

	q := `query Me($id1: string){
		me(func: uid($id1)) {
			id
			name
			age
		}
	}`
	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	//type Root struct {
	//	Me []Person `json:"me"`
	//}

	//var roo Root
	//err = json.Unmarshal(resp.Json, &roo)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(string(resp.Json))
}

func main() {
	route := chi.NewRouter()
	route.Use(middleware.Logger)
	route.Post("/buyers", createBuyer)
	log.Fatal(http.ListenAndServe(":8001", route))
}

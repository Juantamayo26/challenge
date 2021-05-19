package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
)

type db struct {
	Data *allBuyers `json:"all"`
}

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

func GetBuyers(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	txn := dgraphClient.NewTxn()
	ctx := context.Background()
	defer txn.Discard(ctx)

	q := `
    {
			all(func: has(name)) {
				id
				name
				age
			}
    }
	`

	res, err := txn.QueryWithVars(ctx, q, map[string]string{"$a": "Alice"})
	if err != nil {
		fmt.Println(err)
		return
	}
	var data db
	err = json.Unmarshal(res.Json, &data)
	if err != nil {
		log.Println("ERR2")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func CreateBuyer(w http.ResponseWriter, r *http.Request) {
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
	//op := &api.Operation{}
	//op.Schema = `
	//	name: string @index(exact) .
	//	age: int .
	//	married: bool .
	//	loc: geo .
	//`

	ctx := context.Background()
	//err = dg.Alter(ctx, op)
	//if err != nil {
	//	log.Fatal(err)
	//}

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
	log.Println(assigned)

}

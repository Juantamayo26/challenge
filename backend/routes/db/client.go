package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"log"
)

type CancelFunc func()

type Buyer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Person struct {
	Uid string `json:"uid"`
	//Buyer    Buyer   `json:buyer"`
	IP       string    `json:"ip"`
	Device   string    `json:"device"`
	Products []Product `json:"products"`
}

func Query() {
	dg, cancel := getDgraphClient()
	defer cancel()

	p := Person{
		Uid: "_:alice",
		//Buyer: []Buyer{
		//	{id: "0000609f0f00"},
		//},
		IP:     "187.103.198.253",
		Device: "mac",
		Products: []Product{
			{
				ID: "2377d14d",
			},
			{
				ID: "eac57b76",
			},
			{
				ID: "5a40c0a3",
			},
			{
				ID: "de7e99ff",
			},
		},
	}

	op := &api.Operation{}
	op.Schema = `
		Device: string .
		IP: string .

		age: int .
		id string .
		name string .

		type Buyer {
			id
			name
			age
		}

		type Products {
			id
			name
			price
		}
	`

	ctx := context.Background()
	if err := dg.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	//query
	puid := response.Uids["alice"]
	const q = `
		query Me($id: string){
			me(func: uid($id)) {
				ip
			}
		}
	`
	variables := make(map[string]string)
	variables["$id"] = puid
	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	type Root struct {
		Me []Person `json:"me"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)

}

func getDgraphClient() (*dgo.Dgraph, CancelFunc) { // LOGIN
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func ExampleDgraph_Alter_dropAll() {
	dg, cancel := getDgraphClient()
	defer cancel()
	op := api.Operation{DropAll: true}
	ctx := context.Background()
	if err := dg.Alter(ctx, &op); err != nil {
		log.Fatal(err)
	}
	// Output:
}

func main() {
	Query()
}

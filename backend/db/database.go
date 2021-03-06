package db

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Schema() {
	schema := []byte(fmt.Sprintf(`
		type Buyers{
			date: String
			id: String! @id @search(by: [exact])
			name: String
			age: Int
			transaction: [Transactions] @hasInverse(field: buyerid)
		}

		type Products{
			date: String
			id: String! @id
			name: String
			price: Int
			transaction: [Transactions] @hasInverse(field: productids)
		}

		type Transactions{
			date: String
			id: String! @id
			buyerid: Buyers 
			ip: String @search(by: [exact])
			device: String
			productids: [Products]
		}
	`))

	url := "http://localhost:8080/admin/schema"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(schema))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Schema Added")
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func Add(mutation []byte) {
	url := "http://localhost:8080/graphql"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mutation))
	req.Header.Set("Content-Type", "application/graphql")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

package db

import (
	"bytes"
	"fmt"
	//"io/ioutil"
	"net/http"
)

func Schema() {
	schema := []byte(fmt.Sprintf(`
    type Buyers{
      id: String! @id
      name: String
      age: Int
    }

    type Products{
      id: String! @id
      name: String
      price: String
    }

    type Transactions{
      id: String! @id
      buyerid: [Buyers]
      ip: String
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

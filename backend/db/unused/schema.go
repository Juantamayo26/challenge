package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {

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

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

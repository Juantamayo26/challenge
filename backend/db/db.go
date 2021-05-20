package db

import (
	"bytes"
	"net/http"
)

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

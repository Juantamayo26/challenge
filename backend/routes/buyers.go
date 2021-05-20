package routes

import (
	"challenge/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

// Types
type buyer struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func removeDuplicates(arr []string) []string {
	words_string := map[string]bool{}
	for i := range arr {
		words_string[arr[i]] = true
	}
	desired_output := []string{} // Keep all keys from the map into a slice.
	for j, _ := range words_string {
		desired_output = append(desired_output, j)
	}
	return desired_output
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
		if err == io.EOF {
			break
		}

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

	//Adding data to db
	var input []string
	for i := 0; i < len(newBuyer); i++ {
		v := reflect.ValueOf(newBuyer[i])

		temp := []interface{}{"1", "2", 3}
		for i := 0; i < v.NumField(); i++ {
			temp[i%3] = v.Field(i)
			if (i+1)%3 == 0 {
				cur := fmt.Sprintf(`
					{
						id:%q,
						name:%q,
						age:%d
					}
					`, temp...)
				input = append(input, cur)
			}
		}
	}

	uniqueData := removeDuplicates(input) // Remove the duplicates

	inputString := strings.Join(uniqueData, "") // Convert the []string to string
	mutation := []byte(fmt.Sprintf(`
		mutation {
			addBuyers(input:[
			` + inputString + `
			]){
				buyers{
					id
					name
					age
				}
			}
		}
	`))
	db.Add(mutation)
}

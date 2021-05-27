package routes

import (
	"challenge/db"
	"challenge/helper"
	"encoding/json"
	"fmt"
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

func CreateBuyer(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("data")
	date := r.FormValue("date")
	doc := []buyer{}
	jsonDecoder := json.NewDecoder(file)
	err = jsonDecoder.Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

		temp := []interface{}{date, 1, 2, 3} // Save all the values of the struct
		for i := 0; i < v.NumField(); i++ {
			temp[(i%3)+1] = v.Field(i)
			if (i+1)%3 == 0 {
				cur := fmt.Sprintf(`
					{
						date: %q,
						id:%q,
						name:%q,
						age:%d
					}
					`, temp...)
				input = append(input, cur)
			}
		}
	}

	uniqueData := helper.RemoveDuplicates(input) // Remove the duplicates

	inputString := strings.Join(uniqueData, "") // Convert the []string to string
	mutation := []byte(fmt.Sprintf(`
		mutation {
			addBuyers(input:[
			` + inputString + `
			]){
				buyers{
					date
					id
					name
					age
				}
			}
		}
	`))
	db.Add(mutation)
}

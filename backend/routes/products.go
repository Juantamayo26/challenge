package routes

import (
	"challenge/db"
	"challenge/helper"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Types
type Product struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("data")
	fileName := handler.Filename

	outfile, err := os.Create("./temp/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outfile.Close()
	cpy, err := io.Copy(outfile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(cpy)

	csvFile, err := os.Open("./temp/" + fileName)
	if err != nil {
		fmt.Fprintf(w, "Error opening csvFile")
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = '\''

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	var input []string
	for _, record := range records {
		price, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Printf("i=%d, type: %T\n", price, price)
		}
		cur := fmt.Sprintf(`
			{
				id:%q,
				name:%q,
				price:%d
			}
		`, record[0], record[1], price)
		input = append(input, cur)
	}
	uniqueData := helper.RemoveDuplicates(input) // Remove the duplicates
	inputString := strings.Join(uniqueData, "")  // Convert the []string to string
	mutation := []byte(fmt.Sprintf(`
		mutation {
			addProducts(input: [
			` + inputString + `
			]){
				products{
					id
					name
					price
				}
			}
		}
	`))
	fmt.Println(inputString)
	db.Add(mutation)
}

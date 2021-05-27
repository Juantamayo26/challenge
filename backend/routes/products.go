package routes

import (
	"challenge/db"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	content, handler, err := r.FormFile("data")
	date := r.FormValue("date")
	fileName := handler.Filename

	outfile, err := os.Create("./temp/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outfile.Close()
	cpy, err := io.Copy(outfile, content)
	if err != nil {
		fmt.Println(cpy)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	validator := map[string]bool{}
	for _, record := range records {
		price, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Printf("i=%d, type: %T\n", price, price)
		}
		if validator[record[0]] != true {
			validator[record[0]] = true
			cur := fmt.Sprintf(`
				{
					date: %q,
					id:%q,
					name:%q,
					price:%d
				}
			`, date, record[0], record[1], price)
			input = append(input, cur)
		}
	}
	inputString := strings.Join(input, "") // Convert the []string to string
	mutation := []byte(fmt.Sprintf(`
		mutation {
			addProducts(input: [
			` + inputString + `
			]){
				products{
					date
					id
					name
					price
				}
			}
		}
	`))
	//fmt.Println(inputString)
	db.Add(mutation)
}

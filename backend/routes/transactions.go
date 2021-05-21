package routes

import (
	"challenge/db"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func CreateTransactions(w http.ResponseWriter, r *http.Request) {
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

	content, err := ioutil.ReadFile("./temp/" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	var aux string
	var record [5]string
	var index int

	var input []string
	for i := 0; i < len(content); i++ {
		if content[i] != 0 && content[i] != 35 {
			aux = aux + string(content[i])
		}
		if content[i] == 0 && content[i+1] == 0 {

			var p = 1
			var products []string
			for j := 1; j < len(aux); j++ {
				if aux[j] == ',' || aux[j] == ')' {
					cur := fmt.Sprintf(`
						{
							id:%q
						}
						`, aux[p:j])
					products = append(products, cur)
					p = j + 1
				}
			}
			data := fmt.Sprintf(`
				{
					id: %q
					buyerid: {
						id: %q
					}
					ip: %q
					device: %q
					productids: %s
				}
			`,
				record[0], record[1], record[2], record[3], products)
			input = append(input, data)
			index = 0
			aux = ""
			i++
		} else if content[i] == 0 {
			record[index] = aux

			index++
			aux = ""
		}
	}

	fmt.Println(len(input))
	mutation := []byte(fmt.Sprintf(`
		mutation {
			addTransactions(input:
			%s
			){
				transactions{
					id
					buyerid{
						id
					}
					ip
					device
					productids{
						id
					}
				}
			}
		}
	`, input))
	db.Add(mutation)
}

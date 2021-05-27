package routes

import (
	"challenge/db"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("data")
	date := r.FormValue("date")
	fileName := handler.Filename

	outfile, err := os.Create("./temp/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outfile.Close()
	cpy, err := io.Copy(outfile, file)
	if err != nil {
		fmt.Println(cpy)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content, err := ioutil.ReadFile("./temp/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(content), "\000")
	var input []string
	for i := 5; i < len(s); i += 6 {
		x := 6 * (i / 6)
		id := (s[x])
		buyerid := (s[1+x])
		ip := (s[2+x])
		device := (s[3+x])
		productids := (s[4+x])

		var products []string
		var p = 1
		for j := 1; j < len(productids); j++ {
			if productids[j] == ',' || productids[j] == ')' {
				cur := fmt.Sprintf(`
					{
						id:%q
					}`, productids[p:j])
				products = append(products, cur)
				p = j + 1
			}
		}

		data := fmt.Sprintf(`
			{
				date: %q
				id: %q
				buyerid: {
					id: %q
				}
				ip: %q
				device: %q
				productids: %s
			}
		`, date, id, buyerid, ip, device, products)
		input = append(input, data)
	}

	mutation := []byte(fmt.Sprintf(`
		mutation {
			addTransactions(input:
			%s
			){
				transactions{
					date
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

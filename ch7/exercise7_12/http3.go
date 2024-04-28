// Exercise 7.12: Change the handler for /list to print its output as an HTML table, not text.
// You may find html/template package (§4.6) useful.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	const templ = `
	<table>
		<tr>
			<th>Item</th>
			<th>Price</th>
		</tr>
		{{range $key, $value := .}}
		<tr>
			<td>{{$key}}</td>
			<td>{{$value}}</td>
		</tr>
		{{end}}
	</table>
	`

	list, err := template.New("list").Parse(templ)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "unexpected error: %v\n", err)
		return
	}

	list.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

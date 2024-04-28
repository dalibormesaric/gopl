package main

import (
	"fmt"
	"log"
	"main/params"
	"net/http"
)

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe("localhost:12345", nil))
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}
	data.MaxResult = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

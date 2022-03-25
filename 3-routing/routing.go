package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		page := vars["page"]
		book := vars["title"]

		fmt.Fprintf(w, "You've requested the book: %s and the page %s\n", book, page)

	})

	http.ListenAndServe(":83", r)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h1> Welcome to my website </h1> <h2> Picture of the Day </h2> <img src='static/foto.jpeg'>")

	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":83", nil)
}

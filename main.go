package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	mux.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the user page!")
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

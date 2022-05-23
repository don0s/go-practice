package webserver

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func StartSimpleServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<!DOCTYPE html><html><body>Hello!</body></html>")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

package webserver

// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

// func echoString(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello")
// }

func incrementCounter(w http.ResponseWriter, r *http.Request) {

	url := r.URL
	values := url.Query()

	increment := 1
	for k, v := range values {
		fmt.Println(k, " ", v[0])

		if k == "counter" {
			//increment, _ = strconv.Atoi(v[0])
			var err error
			increment, err = strconv.Atoi(v[0])
			fmt.Println(err)
		}
	}

	mutex.Lock()
	counter += increment

	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}

// func serveFavIcon(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "../static/images/favicon.ico")
// }

func StartServerMutex() {

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/images/favicon.ico")
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })

	http.Handle("/", http.FileServer(http.Dir("../static")))

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<!DOCTYPE html><html><body>Hello!</body></html>")
	})

	//log.Fatal(http.ListenAndServe(":8081", nil))
	//log.Fatal(http.ListenAndServeTLS(":8443", "../server.crt", "../server.key", nil))
	startListening()
}

func startListening() {

	go func() {
		err_http := http.ListenAndServe(":8081", nil)
		if err_http != nil {
			log.Fatal("Web Server (HTTP): ", err_http)
		}
	}()

	err_http := http.ListenAndServeTLS(":8443", "../server.crt", "../server.key", nil)
	if err_http != nil {
		log.Fatal("Web Server (HTTPS): ", err_http)
	}
}

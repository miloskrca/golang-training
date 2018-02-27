package main

import (
	"log"
	"net/http"
)

// http://localhost:8080/demo/http/fileserver/fileserver.go
func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

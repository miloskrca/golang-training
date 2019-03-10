package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	http.HandleFunc("/ciao", ciao)
	http.Handle("/holla", &Holla{"Everybody"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Holla is our handler struct
type Holla struct {
	defaultValue string
}

func (h *Holla) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	to := r.URL.Query().Get("to")
	if to == "" {
		to = h.defaultValue
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello, %s", to))
}

func ciao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ciao")
}

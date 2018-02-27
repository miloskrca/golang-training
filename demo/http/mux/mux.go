package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", index)
	router.Handle("/user/", newUserRouter())
	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}

func newUserRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", user)
	return router
}

func user(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/user/")
	w.Write([]byte("User " + user))
}

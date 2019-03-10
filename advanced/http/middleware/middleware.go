package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.Handle("/index/", newRouter())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", logRequest(stripPath("/index/", http.HandlerFunc(index))))
	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "URL ", r.URL.Path)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Println("Method:", r.Method, "Path:", r.URL.Path, "Took:", t2.Sub(t1))
	})
}

func stripPath(subpath string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, subpath)
		r.URL.Path = path
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type key string

var userAgentKey = key("User-Agent")

func main() {
	http.Handle("/", userAgent(http.HandlerFunc(index)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func userAgent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), userAgentKey, r.Header.Get("User-Agent"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	userAgent, ok := r.Context().Value(userAgentKey).(string)
	if !ok {
		log.Fatal("userAgent not a string")
	}
	fmt.Fprint(w, "User-Agent: ", userAgent)
}

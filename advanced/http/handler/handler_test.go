package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCiao(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/ciao", nil)
	w := httptest.NewRecorder()
	ciao(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Ciao endpoint did not return 200. Returned: %d", w.Code)
	}
}

func TestHolla(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/holla?to=Me", nil)
	w := httptest.NewRecorder()
	holla := &Holla{"Everybody"}
	holla.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Holla endpoint did not return 200. Returned: %d", w.Code)
	}
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != "Hello, Me" {
		t.Errorf("Holla endpoint did not return 'Hello Me'. Returned: %s", body)
	}
}

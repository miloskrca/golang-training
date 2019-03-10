package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// RenderJSON render a content as json(thinking about middleware)
func RenderJSON(w http.ResponseWriter, content interface{}, statusCode int) {
	// Set Content-Type as json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// HTTP STATUS CODE
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		log.Println(err)
	}
}

type Card struct {
	Name       string
	Number     string
	CVC        int
	Type       string `json:"CardType"`
	updatedAt  time.Time
	VerifiedBy string `json:",omitempty"`
}

func (c Card) Validate() error { return nil }
func (c Card) Save() {
	fmt.Printf("%+v\n", c)
}

func createCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderJSON(w, map[string]string{"errors": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	var card Card
	err := json.NewDecoder(r.Body).Decode(&card)
	defer r.Body.Close()
	if err != nil {
		RenderJSON(w, map[string]string{"errors": err.Error()}, http.StatusInternalServerError)
		return
	}
	if err := card.Validate(); err != nil {
		RenderJSON(w, err, http.StatusBadRequest)
	}
	card.updatedAt = time.Now()
	card.Save()
	RenderJSON(w, card, http.StatusCreated)
}

func main() {
	http.HandleFunc("/card", createCard)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -H "Content-Type: application/json" -X POST -d '{"Name": "Card Holder", "Number": "424242424242", "CVC": 999, "CardType": "Visa"}' http://localhost:8080/card

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota // 0
	Gopher                // 1
	Zebra                 // 2
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

func main() {
	blob := []byte(`["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`)
	var zoo []Animal
	if err := json.Unmarshal(blob, &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal]++
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}

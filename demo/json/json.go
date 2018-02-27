package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	fmt.Println("====================")
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	fmt.Println("====================")

	var msg Message
	if err := json.Unmarshal(b, &msg); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	fmt.Println("====================")

	mymap := f.(map[string]interface{})
	for k, v := range mymap {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	fmt.Println("====================")
}

package main

import (
	"errors"
	"fmt"
)

func doSomething() (string, error) {
	return "str", errors.New("error")
}

func main() {
	if str, err := doSomething(); err != nil {
		panic(err)
	} else {
		fmt.Printf("String: %s", str)
	}
}

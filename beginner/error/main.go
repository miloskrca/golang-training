package main

import (
	"errors"
)

func doSomething() error {
	return errors.New("error")
}

func main() {
	err := doSomething()
	if err != nil {
		panic(err)
	}
}

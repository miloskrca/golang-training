package main

import (
	"fmt"
)

func main() {
	var i int
	go func() {
		i = 5
	}()
	// time.Sleep(1 * time.Second)
	fmt.Println("i", i)
}

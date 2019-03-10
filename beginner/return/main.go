package main

// import required modules
import (
	"fmt"
)

// "named" return
func named() (str string) {
	str = "lorem"
	return
}

// "typed" return
func typed() string {
	var str = "ipsum"
	return str
}

// main function
func main() {
	fmt.Println(named())
	fmt.Println(typed())
}

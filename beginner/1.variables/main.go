package main

// import required modules
import (
	"fmt"
)

// in go you can write comments after "//"

/*
to write multiline comments you can write
inside of "/ *" and "* /" (without the spaces
*/

func main() {
	// declare a variable
	var str = "Hello World"

	// declare a variable short syntax
	str2 := "Hello Planet"

	// declare uninitalized variable
	var str3 string

	// print variables
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str3)
}

package main

import "fmt"

// Variadic functions can have zero or n parameters passed
// The arguments passed to a variadic function are appended to a slice of the same type
func variadic(numbers ...int) {
	fmt.Printf("Type: %T\t Content: %d\n", numbers, numbers)
	// a variadic parameter is accessed as an array inside the function
	for _, number := range numbers {
		fmt.Printf("Number: %d\n", number)
	}
}

func main() {
	variadic(1, 2, 3, 4, 5)
}

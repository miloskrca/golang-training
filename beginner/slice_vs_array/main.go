package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	// basic
	a := [3]int{1, 2, 3} // Array literal
	b := a               // Copy the contents of a into b
	a[0] = 0
	fmt.Println(a) // Prints "[0 2 3]"
	fmt.Println(b) // Prints "[1 2 3]"
	fmt.Println("=============")

	// difference when updating slices vs arrays
	var array = [2]point{
		point{1, 2},
		point{3, 4},
	}
	var slice = array[:1]
	slice[0].x = 111
	array[0].y = 222

	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println("=============")

	changer(array)
	fmt.Println(array)
	fmt.Println("=============")
}

func changer(array [2]point) {
	array[0].x = 121212
}

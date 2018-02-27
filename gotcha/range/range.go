package main

import "fmt"

type animal struct {
	name string
	legs int
}

func main() {
	zoo := []animal{animal{"Dog", 4},
		animal{"Chicken", 2},
		animal{"Snail", 0},
	}

	fmt.Printf("-> Before update %v\n", zoo)

	for _, animal := range zoo {
		// 🚨 Oppps! `animal` is a copy of an element 😧
		animal.legs = 999
	}

	fmt.Printf("\n-> After update %v\n", zoo)
}

// for idx, _ := range zoo {
// 	zoo[idx].legs = 999
// }

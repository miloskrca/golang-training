package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	// expecting to print all numbers 0..9
	for i := 0; i <= 9; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

// the runtime scheduler is optimizing the code,
// so the goroutines are not blocked when executing the logic and it schedules IO for later
// i := i

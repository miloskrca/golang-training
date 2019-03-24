package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(3 * time.Second)
		}
	}()

	timeout := time.After(6 * time.Second)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Message 1", msg1)
		case msg2 := <-c2:
			fmt.Println("Message 2", msg2)
		// case <-time.After(6 * time.Second):
		case <-timeout:
			fmt.Println("stopping after 6 seconds")
			return
		}
	}
}

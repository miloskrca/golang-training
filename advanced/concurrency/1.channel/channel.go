package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)

	go pinger(c)
	go printer(c)

	fmt.Scanln()
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	stop := make(chan struct{})
	done := make(chan struct{})

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

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-stop:
				fmt.Println("stopping after manual action")
				close(done)
			}
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		s := <-c
		fmt.Println("received signal:", s)
		stop <- struct{}{}
	}()
	<-done
	fmt.Println("done")
}

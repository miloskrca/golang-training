package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	go do(ctx, done)
	<-done
	fmt.Println("Done")
}

func do(ctx context.Context, done chan struct{}) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Done after 3 seconds")
	// If the context finishes before we could get the result, exit early
	case <-ctx.Done():
		fmt.Println("Context done after cancelation")
	}
	done <- struct{}{}
}

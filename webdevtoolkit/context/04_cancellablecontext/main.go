package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := range gen(ctx) {
		fmt.Println(i)
		if i == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

// leaky gen
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

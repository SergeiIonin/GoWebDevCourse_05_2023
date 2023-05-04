package main

import "fmt"

func main() {

	for i := range gen() {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

}

// leaky gen
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}

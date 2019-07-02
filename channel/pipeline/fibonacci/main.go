package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan<- int, quit chan struct{}) {
	x, y := 0, 1
	defer close(c)
	for {
		select {
		case c <- x:
			x, y = y, x+y
			time.Sleep(time.Millisecond * 1000)
		case <-quit:
			fmt.Println("quit")
			// exit for loop and exit fib func
			return
		}
	}
}

func main() {
	// fib elements
	c := make(chan int)
	// control channel
	quit := make(chan struct{})
	go func() {
		for i := 0; i < 3; i++ {
			// wait until c channel have a element
			fmt.Println(<-c)
		}
		close(quit)
	}()
	fibonacci(c, quit)
}

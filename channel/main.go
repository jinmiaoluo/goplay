package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() { ch <- 1 }()
	go func() { ch <- 2 }()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

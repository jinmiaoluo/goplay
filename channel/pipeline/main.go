package main

import "fmt"

// The first stage, gen, is a function that converts a list of integers to a
// channel that emits the integers in the list. The gen function starts a
// goroutine that sends the integers on the channel and closes the channel when
// all the values have been sent:
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// The second stage, sq, receives integers from a channel and returns a channel
// that emits the square of each received integer. After the inbound channel is
// closed and this stage has sent all the values downstream, it closes the
// outbound channel:
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// The main function sets up the pipeline and runs the final stage: it receives
// values from the second stage and prints each one, until the channel is
// closed:
//func main() {
//	// Set up the pipeline.
//	c := gen(2, 3)
//	out := sq(c)
//
//	// Consume the output.
//	fmt.Println(<-out) // 4
//	fmt.Println(<-out) // 9
//}
func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

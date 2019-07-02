package main

import (
	"fmt"
	"sync"
)

// The first stage, gen, is a function that converts a list of integers to a
// channel that emits the integers in the list. The gen function starts a
// goroutine that sends the integers on the channel and closes the channel when
// all the values have been sent:
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
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

// The merge function converts a list of channels to a single channel by
// starting a goroutine for each inbound channel that copies the values to the
// sole outbound channel. Once all the output goroutines have been started,
// merge starts one more goroutine to close the outbound channel after all
// sends on that channel are done.
//
// Sends on a closed channel panic, so it's important to ensure all sends are
// done before calling close. The sync.WaitGroup type provides a simple way to
// arrange this synchronization:
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// A WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of
	// goroutines to wait for. Then each of the goroutines
	// runs and calls Done when finished. At the same time,
	// Wait can be used to block until all goroutines have finished.
	//
	// A WaitGroup must not be copied after first use.
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	// Add adds delta, which may be negative, to the WaitGroup counter.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics.
	//
	// Note that calls with a positive delta that occur when the counter is zero
	// must happen before a Wait. Calls with a negative delta, or calls with a
	// positive delta that start when the counter is greater than zero, may
	// happen at any time.
	// Typically this means the calls to Add should execute before the statement
	// creating the goroutine or other event to be waited for.
	// If a WaitGroup is reused to wait for several independent sets of events,
	// new Add calls must happen after all previous Wait calls have returned.
	// See the WaitGroup example.
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Multiple functions can read from the same channel until that channel is
// closed; this is called fan-out. This provides a way to distribute work
// amongst a group of workers to parallelize CPU use and I/O.

// A function can read from multiple inputs and proceed until all are closed by
// multiplexing the input channels onto a single channel that's closed when all
// the inputs are closed. This is called fan-in.
//
// We can change our pipeline to run two instances of sq, each reading from the
// same input channel. We introduce a new function, merge, to fan in the
// results:
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2)
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}
}

// goroutine can execute at anytime before or after other function.
package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hi")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

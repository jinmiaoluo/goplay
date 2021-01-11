package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func catchUSR1Signal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGUSR1:
			fmt.Println("USR1 signal catched")
		}
	}
}

func main() {
	catchUSR1Signal()
}

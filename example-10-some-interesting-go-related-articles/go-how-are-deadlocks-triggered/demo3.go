package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	c := make(chan bool)

	select {
	case <-c:
	case <-s:
		println("program stopped")
	}

	fmt.Println("exit")
}

package main

import (
	"log"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
)

func main() {
	f, err := os.OpenFile("trace.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)

	c := make(chan bool)

	select {
	case <-c:
	case <-s:
		trace.Stop()
		println("program stopped")
	}
}

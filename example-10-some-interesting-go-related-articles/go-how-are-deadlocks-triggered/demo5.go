package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)

	c := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("in a goroutine")
		<-c
	}()

	select {
	case <-c:
	case <-s:
		println("program stopped")
	}
}

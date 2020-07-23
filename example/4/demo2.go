package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	go a()
	m1()
}
func m1() {
	m2()
}
func m2() {
	m3()
}
func m3() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	time.Sleep(time.Hour)
}
func a() {
	time.Sleep(time.Hour)
}

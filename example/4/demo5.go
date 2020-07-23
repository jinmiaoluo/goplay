package main

import (
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
	panic("panic from m3")
}
func a() {
	time.Sleep(time.Hour)
}

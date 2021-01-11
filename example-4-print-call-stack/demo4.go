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
	time.Sleep(time.Hour)
}
func a() {
	time.Sleep(time.Hour)
}

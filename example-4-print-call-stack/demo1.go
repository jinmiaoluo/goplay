package main

import (
	"fmt"
	"runtime/debug"
)

func doo() {
	fmt.Println("Here is function doo")
	debug.PrintStack()
}

func coo() {
	doo()
}

func boo() {
	coo()
}

func main() {
	boo()
}

package main

import (
	"fmt"
)

const (
	a int = 0
	b int = 1 << iota
	c
	d
)

func main() {
	fmt.Printf("a = %+016b\n", a)
	fmt.Printf("b = %+016b\n", b)
	fmt.Printf("c = %+016b\n", c)
	fmt.Printf("d = %+016b\n", d)
}

package main

import (
	"fmt"
	"unsafe"
)

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func main() {
	var a float64 = 3.00000
	b := Float64bits(a)
	fmt.Println(b)
}

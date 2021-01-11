package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	var i uintptr = 1
	// equivalent to e := unsafe.Pointer(&x[i])
	g := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0])))
	fmt.Println(g) // 2
}

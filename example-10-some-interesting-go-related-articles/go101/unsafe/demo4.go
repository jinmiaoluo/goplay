package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//a := make([]int, 5)
	//for i := 0; i < len(a); i++ {
	//	a[i] = i
	//}
	//fmt.Println(a)

	// equivalent to e := unsafe.Pointer(&x[i])
	//e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))

	var s struct {
		f int
		i string
	}

	s = struct {
		f int
		i string
	}{f: 1, i: "hello" }

	// equivalent to f := unsafe.Pointer(&s.f)
	f := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f)))
	i := *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.i)))

	fmt.Println(f) // 1
	fmt.Println(i) // hello
}
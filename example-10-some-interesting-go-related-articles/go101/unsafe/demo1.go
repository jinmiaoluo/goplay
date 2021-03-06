package main

import "fmt"
import "unsafe"

func main() {
	var x struct {
		a int64
		b bool
		c string
	}
	const M, N, O, P= unsafe.Sizeof(x.a), unsafe.Sizeof(x.b), unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(M, N, O, P) // 16 32

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16
}

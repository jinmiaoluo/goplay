package main

import (
	"fmt"
)
import "unsafe"

func main() {
	type T struct {
		c string
	}
	type S struct {
		b bool
	}
	var x struct {
		a int64
		*S
		T
	}

	fmt.Println(unsafe.Offsetof(x.a)) // 0

	fmt.Println(unsafe.Offsetof(x.S)) // 8
	fmt.Println(unsafe.Offsetof(x.T)) // 16

	// 此行可以编译过，因为选择器x.c中的隐含字段T为非指针。
	fmt.Println(unsafe.Offsetof(x.c)) // 16

	// 此行编译不过，因为选择器x.b中的隐含字段S为指针。
	//fmt.Println(unsafe.Offsetof(x.b)) // error

	// 此行可以编译过，但是它将打印出字段b在x.S中的偏移量.
	fmt.Println(unsafe.Offsetof(x.S.b)) // 0
}

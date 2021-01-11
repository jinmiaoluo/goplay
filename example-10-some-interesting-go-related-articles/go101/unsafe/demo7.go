package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := [6]byte{'G', 'o', '1', '0', '1'}
	bs := []byte("Golang")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	hdr.Data = uintptr(unsafe.Pointer(&a))

	hdr.Len = 2
	hdr.Cap = len(a)
	fmt.Printf("%s\n", bs) // Go
	bs = bs[:cap(bs)]
	fmt.Printf("%s\n", bs) // Go101
}

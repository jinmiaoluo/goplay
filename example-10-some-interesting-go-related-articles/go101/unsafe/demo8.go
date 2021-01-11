package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// 传递给此函数的string实参在此函数退出后不应被继续使用。
func String2ByteSlice(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Cap = strHdr.Len
	sliceHdr.Len = strHdr.Len
	return
}

func main() {
	// str := "Golang"
	// 对于官方标准编译器来说，上面这行将使str中的字节
	// 开辟在不可修改内存区。所以这里我们使用下面这行。
	str := strings.Join([]string{"Go", "land"}, "")
	s := String2ByteSlice(str)
	fmt.Printf("%s\n", s) // Goland
	s[5] = 'g'
	fmt.Println(str) // Golang
}

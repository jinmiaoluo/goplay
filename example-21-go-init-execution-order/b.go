package main

import (
	"fmt"

	"github.com/jinmiaoluo/goplay/example/21/d"
)

func init() {
	fmt.Println("### in b.go file init function")
	fmt.Println("d.Message ->", d.Message)
	fmt.Println("change d.Message value")
	d.Message = "b"
	fmt.Println("d.Message ->", d.Message)
}

func sayHi() {
	fmt.Println("sayHi function from b.go respone: Hi!")
}

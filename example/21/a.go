package main

import (
	"fmt"

	"github.com/jinmiaoluo/goplay/example/21/d"
)

func init() {
	fmt.Println("### in a.go file init function")
	fmt.Println("d.Message ->", d.Message)
	fmt.Println("change d.Message value")
	d.Message = "a"
	fmt.Println("d.Message ->", d.Message)
}

func init() {
	fmt.Println("### in a.go file the second init function")
	fmt.Println("d.Message ->", d.Message)
}

func main() {
	fmt.Println("### in a.go file main function")
	fmt.Println("d.Message ->", d.Message)
	fmt.Println("call function from b.go")
	sayHi()
}

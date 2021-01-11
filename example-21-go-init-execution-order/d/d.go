package d

import "fmt"

var Message string

func init() {
	fmt.Println("### in d.go file init function")
	fmt.Println("set default d.Message value")
	Message = "d"
	fmt.Println("d.Message ->", Message)
}

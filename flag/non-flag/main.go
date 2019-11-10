package main

import (
	"flag"
	"fmt"
)

func main() {
	v := flag.Int("num", 0, "a number")
	s := flag.String("str", "string", "a string")
	flag.Parse()
	fmt.Println(int(*v))
	fmt.Println(string(*s))
	// go run main.go  -- hello world  again
	fmt.Println(flag.Args()) // [hello world again]
	fmt.Println(flag.NArg()) // 3
	fmt.Println(flag.Arg(0)) // hello
}

// Package main provides ...
package main

import (
	"flag"
	"fmt"
)

// redefined flag means more than one flag use the same name.
func main() {
	a := flag.Bool("bool", false, "test redefined bool flag")
	b := flag.Bool("bool", false, "test redefined bool flag")
	//flag.CommandLine.Bool("bool", false, "test redefined bool flag")
	//flag.CommandLine.Bool("bool", false, "test redefined bool flag")
	fmt.Println(a)
	fmt.Println(b)
	flag.Parse()
}

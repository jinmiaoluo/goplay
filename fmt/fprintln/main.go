package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hello world")
	fmt.Fprintln(os.Stderr, "hello world with error")
}

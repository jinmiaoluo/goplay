// Package main provides ...
package main

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
)

func main() {
	fmt.Println(homedir.Dir())
	fmt.Println(homedir.Expand("~/go/src"))
}

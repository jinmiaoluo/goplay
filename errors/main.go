// Package main provides ...
package main

import (
	"errors"
	"fmt"
)

func main() {
	testErr := errors.New("test error from function")
	fmt.Println(testErr)
}

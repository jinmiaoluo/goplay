package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == i {
				fmt.Printf("%dX%d=%d\n", j, i, i*j)
				break
			}
			fmt.Printf("%dX%d=%d\t", j, i, i*j)
		}
	}
}

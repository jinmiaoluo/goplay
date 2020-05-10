package main

import "fmt"

func main() {
	a := []int{91, 83, 105, 109, 112, 108, 105, 102, 105, 101, 100, 32, 66, 83, 68, 32, 76, 105, 99, 101, 110, 115, 101, 93, 40, 76, 73, 67, 69, 78, 83, 69, 46, 116, 120, 116, 41}
	var b string
	for i := 0; i < len(a); i++ {
		b += string(a[i])
	}
	fmt.Println(b)
}

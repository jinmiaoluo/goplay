// convert string into a []int
package main

import "fmt"

func main() {
	a := "hello world"
	b := []int{}
	for i := 0; i < len(a); i++ {
		b = append(b, int(a[i]))
	}
	fmt.Println(b)
}

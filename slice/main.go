package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("before delete: ", a)
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == 3 {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}
	fmt.Println("after delete: ", a)
}

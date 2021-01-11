package main

import (
	"fmt"
	"sort"
)

var a []int = []int{3, 5, 2, 8, 9, 12, 99, 23, 1, 44, 22, 33, 34, 32, 12, 55, 4, 1}

func main() {
	sort.IntSlice(a).Sort()
	fmt.Println(a)
}

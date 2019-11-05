package main

import "fmt"

var b = []int{1, 2, 3, 4, 5, 6, 7}

func bs(b []int, v int) int {
	i := 0
	j := len(b)
	for i < j {
		m := i + (j-i)/2
		b_ := b[m]
		if v == b_ {
			return b_
		} else if v > b_ {
			i = m + 1
		} else if v < b_ {
			j = m - 1
		}
	}
	return -1
}

func main() {
	v := bs(b, 2)
	fmt.Println(v)
}

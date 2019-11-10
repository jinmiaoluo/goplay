package main

import "fmt"

// fibonacci implementation without Closure
func fibonacci(num int) []int {
	fs := make([]int, num, num)
	a, b := 0, 1
	for i := 0; i < num; i++ {
		fs[i] = a
		a, b = b, a+b
	}
	return fs
}
func main() {
	fmt.Println(fibonacci(10))
}

/*
// fibonacci implementation with Closure
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
*/

/*
// fibonacci implementation with Recursion
func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num <= 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	fmt.Println(fibonacci(10))
}
*/

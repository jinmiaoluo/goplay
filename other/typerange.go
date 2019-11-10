package main

import "fmt"

const IMin int64 = -1 << 63
const IMax int64 = 1<<63 - 1
const JMin int8 = -1 << 7
const JMax int64 = 1<<7 - 1

func main() {
	fmt.Println(IMin)
	fmt.Println(IMax)
	fmt.Println(JMin)
	fmt.Println(JMax)
}

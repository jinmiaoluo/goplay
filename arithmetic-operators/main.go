package main

import "fmt"

func main() {
	/*
		And Not ( bit clear )
		位清除: 将指定位的值置0
	*/

	fmt.Printf("%08b\n", 0b11111111&^0b00000001)
	// 将 `11111111` 的最后1位置零
	// Output: 11111110

	fmt.Printf("%08b\n", 0b11111111&^0b11110000)
	// 将 `11111111` 的前4位置零
	// Output: 00001111
}

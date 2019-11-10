package main

import "fmt"

func main() {
	aMap := make(map[string]interface{})
	aMap["i1"] = 1
	aMap["i2"] = 2

	var v interface{}
	var ok bool

	v, ok = aMap["i1"]
	fmt.Println(v, ok)
	v, ok = aMap["i3"]
	fmt.Println(v, ok)
}

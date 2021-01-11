package main

import (
	"github.com/tidwall/randjson"
)

func main() {
	js := randjson.Make(12, nil)
	println(string(js))
}

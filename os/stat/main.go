package main

import (
	"fmt"
	"os"
)

func main() {
	i, _ := os.Stat("/tmp")
	fmt.Println(i.Mode(), i.Name(), i.Size(), i.Sys(), i.ModTime(), i.IsDir())

	_, e := os.Stat("127.0.0.1")
	fmt.Println(e.Error())
}

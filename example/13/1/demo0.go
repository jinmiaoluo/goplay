package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Stat("small.txt")
	if err != nil {
		if os.IsNotExist(err) { // 判断文件是否存在
			log.Fatal("File isn't existed")
		} else {
			log.Fatal(err)
		}
	}
	Fsize := f.Size()
	fmt.Printf("size is: %d bytes\n", Fsize)
}

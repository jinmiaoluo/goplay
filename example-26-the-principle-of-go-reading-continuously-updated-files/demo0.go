package main

import (
	"fmt"

	"github.com/nxadm/tail"
)

func main() {
	t, err := tail.TailFile("./demo.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

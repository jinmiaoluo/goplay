// Package main provides ...
package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip = flag.Int("ip", 1234, "help message for flagname")
	var hostname = flag.String("hostname", "jinmiaoluo.cn", "hostname for connection")
	flag.Parse()
	fmt.Println(*ip)
	fmt.Println(*hostname)
}

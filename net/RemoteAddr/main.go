package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "master.beansmile-dev.com:443")
	fmt.Printf("Conn's remote address is %s\n", conn.RemoteAddr().String())
}

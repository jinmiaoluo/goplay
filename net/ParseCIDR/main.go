package main

import (
	"fmt"
	"log"
	"net"
)

func ExampleParseCIDR() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)

	// Output:
	// 192.0.2.1
	// 192.0.2.0/24
	// 2001:db8:a0b:12f0::1
	// 2001:db8::/32
}

func main() {
	ExampleParseCIDR()
}

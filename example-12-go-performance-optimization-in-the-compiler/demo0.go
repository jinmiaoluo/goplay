package main

import (
	"fmt"
	"strings"
	"testing"
)

var gogogo = strings.Repeat("Go", 1024)

func f() {
	for range []byte(gogogo) {
	}
}

func m() {
	b := []byte(gogogo)
	for range b {
	}
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, m))
}

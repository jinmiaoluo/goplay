package demo

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before test")
	ret := m.Run()
	fmt.Println("after test")
	os.Exit(ret)
}

func TestA(t *testing.T) {
	fmt.Println("test for A")
}

func TestB(t *testing.T) {
	fmt.Println("test for B")
}

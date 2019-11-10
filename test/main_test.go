package main

import "testing"

func TestprintHello(t *testing.T) {
	result := printHello("Hello")
	if result != "Hello" {
		t.Errorf("expecting Hello, got %s", result)
	}
}

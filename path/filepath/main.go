package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if currentWorkingDir, err := os.Getwd(); err == nil {
		fmt.Println(filepath.Join(currentWorkingDir, "config.json"))
	}
}

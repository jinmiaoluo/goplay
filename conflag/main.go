// Package  provides ...
package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/monochromegane/conflag"
)

func main() {
	// define your flags.
	var procs int
	configPath, _ := homedir.Expand(".conflag.toml")
	flag.IntVar(&procs, "procs", runtime.NumCPU(), "GOMAXPROCS")

	// set flags from configuration before parse command-line flags.
	if args, err := conflag.ArgsFrom(configPath); err == nil {
		_ = flag.CommandLine.Parse(args)
	}

	// parse command-line flags.
	flag.Parse()
	fmt.Println("the value of procs is:", procs)
}

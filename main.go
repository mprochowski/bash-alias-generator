package main

import (
	"github.com/mprochowski/bash-alias-generator/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Run())
}

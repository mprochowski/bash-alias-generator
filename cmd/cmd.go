package cmd

import (
	"flag"
	"fmt"
	"github.com/mprochowski/bash-alias-generator/pkg/generator"
	"io/ioutil"
	"os"
)

func Run() int {
	p := flag.String("f", "", "Json/Yaml definition file.")
	flag.Parse()

	if *p == "" {
		fmt.Fprintf(os.Stderr, "Error: Invalid file path.\n")
		return 1
	}

	d, err := ioutil.ReadFile(*p)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v.\n", err)
		return 1
	}

	r, err := generator.Generate(string(d))
	for _, v := range r {
		fmt.Println(v)
	}

	return 0
}

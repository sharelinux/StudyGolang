package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 帮助方法1
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "qusting")
		flag.PrintDefaults()
	}

	flag.Parse()
	fmt.Printf("Hello. %s\n", name)
}

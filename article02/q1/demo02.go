package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everynoe", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s\n", name)
}
package main

import (
"flag"
"fmt"
"os"
)

var name string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	cmdLine.StringVar(&name,"name","everyone","the greeting someone")
}

func main() {
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s\n", name)
}
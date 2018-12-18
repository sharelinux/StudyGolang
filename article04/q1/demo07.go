package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string                                                  					  // [1]
	flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)

}
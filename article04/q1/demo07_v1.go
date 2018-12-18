package main

import (
	"flag"
	"fmt"
)

func main() {
	// 方式1
	var name = flag.String("name", "everyone", "The greeting object.")

	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name) // 适用于方式1
}
package main

import (
	"flag"
	"fmt"
)

var Name = flag.String("name", "allen", "Please input your name")
var Age = flag.Int("age", 28, "Please input your age")
var SomeInt int

func Init() {
	flag.IntVar(&SomeInt, "flagname", 8537, "Help message for flagname")
}


func main() {
	Init()
	flag.Parse()

	// After parsing, the arguments after the flag are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1
	// Args returns the non-flag command-line arguments
	// NArg is the number of arguments remaining after flags have been processed
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("args[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Printf("name=%s\n", *Name)
	fmt.Printf("age=%d\n", *Age)
	fmt.Printf("flagname=%d\n", SomeInt)
}
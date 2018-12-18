package main

import (
	"flag"

	"StudyGolang/article03/q3/lib"
	//in "StudyGolang/article03/q3/lib/internal" // 此行无法通过编译。
	//os
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting someone!")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}
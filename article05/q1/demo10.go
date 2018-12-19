package main

import "fmt"

var block = "package"

func main() {
	fmt.Printf("The block is %s.\n", block)

	block := "function"  // 变量重声明
	{
		block := "inner" // 变量重声明
		fmt.Printf("The block is %s.\n", block)
	}

	fmt.Printf("The block is %s.\n", block)
}
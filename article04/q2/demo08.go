package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag() // 可以自动类型推断
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

// 字符串的指针类型
func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting someone.")
}


// int的指针类型
//func getTheFlag() *int {
//	return flag.Int("num", 0, "The number of the luck!")
//}
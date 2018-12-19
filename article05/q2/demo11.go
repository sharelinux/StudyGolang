package main

import "fmt"

var container = []string{"0", "1", "2"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2:"two"} // 可重名变量 屏蔽全局container变量
	container = map[int]string{1: "one", 2: "two", 3:"three"} // 变量重声明
	fmt.Printf("The element is %q.\n", container[1])
}
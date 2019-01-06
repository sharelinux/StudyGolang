package main

import (
	"container/ring"
	"fmt"
)

func main() {
	runRing()
}

func runRing() {
	r := ring.New(10) // 初始化长度为10的环形链表
	for i := 0; i < r.Len(); i++ {
		r.Value = i // 链表赋值
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%v ", r.Value)
		r = r.Next()
	}
	fmt.Println()

	r = r.Move(3) // 向后移动3个
	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%v ", r.Value)
		r = r.Next()
	}
	fmt.Println()

	r1 := r.Unlink(7)
	for i := 0; i < r1.Len(); i++ {
		fmt.Printf("%v ", r1.Value)
		r1 = r1.Next()
	}
	fmt.Println()
	fmt.Println(r1.Len())
}

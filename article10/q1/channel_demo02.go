package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	// 示例1
	ch1 := make(chan []int, 1)
	s1 := []int{1, 2, 3}
	ch1 <- s1
	s2 := <-ch1
	s2[0] = 100
	fmt.Printf("%#v\n%#v\n", s1, s2)
	fmt.Printf("%p\n%p\n", s1, s2)
	fmt.Println()

	// 示例2
	ch2 := make(chan [3]int, 1)
	s11 := [3]int{1, 2, 3}
	ch2 <- s11
	s22 := <-ch2
	s22[0] = 100
	fmt.Printf("%#v\n%#v\n", s11, s22)
	fmt.Printf("%p\n%p\n", s11, s22)
	fmt.Println()

	// 示例3
	ch3 := make(chan *Student, 1)
	stu1 := &Student{"allen", 28}
	ch3 <- stu1
	stu2 := <-ch3

	fmt.Printf("stu1: %p\nstu2: %p\n", stu1, stu2)
}

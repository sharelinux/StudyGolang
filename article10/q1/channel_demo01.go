package main

import "fmt"

func main() {
	// 示例1 非缓冲通道 长度=容量
	ch1 := make(chan int)
	fmt.Printf("ch1 len(%v), cap(%v)\n", len(ch1), cap(ch1))

	// 示例2 缓冲通道 len代表通道中已存入的元素数量，cap代表整个缓冲区的大小
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	fmt.Printf("ch2 len(%v), cap(%v)\n", len(ch2), cap(ch2))

	// 示例3 缓冲通道 当缓冲通道已满的时候，长度=容量
	ch2 <- 3
	fmt.Printf("ch2 len(%v), cap(%v)\n", len(ch2), cap(ch2))
}

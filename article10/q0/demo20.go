package main

import "fmt"

func main() {
	// 示例1
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
	fmt.Println()

	// 示例2
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)

	//如果上面的ch2通道不关闭，那么range函数就不会结束，从而在接收第四个数据的时候就阻塞了。
	for elem := range ch2 {
		fmt.Printf("channel ch2: %v\n", elem)
	}
	fmt.Println()

	// 示例3
	ch3 := make(chan int, 3)
	ch3 <- 3
	ch3 <- 2
	ch3 <- 1
	close(ch3) // 必须关闭
	for {
		if v, ok := <-ch3; ok {
			fmt.Printf("channel ch3: %v\n", v)
		} else {
			break // 表示channel已经被关闭，退出循环
		}

	}

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()

}

func example1() {
	// 准备几个通道
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	// 随机选择一个通道， 并向它发送元素.
	rand.Seed(time.Now().UnixNano()) // 随机数种子
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index

	// 哪一个通道有可取的元素之，哪个对应的分支就会被执行.
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate is selected.")
	case <-intChannels[2]:
		fmt.Println("The third candidate is selected.")
	default:
		fmt.Println("No candidate case is selected.")
	}
}

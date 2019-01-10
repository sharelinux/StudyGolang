package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 创建一个channel，用于goroutine之间通知情况
	ch1 := make(chan bool, 1)
	// 开启goroutine，并把ch1 channel传进去
	go example_goroutines(ch1)

	//如果ch1 channel中一直没有数据，这里就会卡住，直到example_goroutines结束时传入值以后，这里才会继续执行
	<-ch1

}

func example_goroutines(ch chan<- bool) { //函数参数为发送的单项通道
	// 准备几个通道
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	// 随机选择一个通道， 并向它发送元素.
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
	// 向channel中传入值true
	ch <- true
	close(ch)
}

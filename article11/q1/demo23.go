package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 示例1
	// 只能发, 不能收的通道
	var uselessChan = make(chan<- int, 1)
	// 只能收, 不能发送的通道
	var anotherUserlessChan = make(<-chan int, 1)
	// 这里分别打印代表2个通道的16进制表示
	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan, anotherUserlessChan)
	fmt.Println()

	// 示例2
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	// 示例4
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}
	fmt.Println()

	// 示例5
	intChan3 := GetIntChan(getIntChan)
	for elem := range intChan3() {
		fmt.Printf("The element in intChan3: %v\n", elem)
	}
}

// 示例2
// 只能发送的channel方法
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
	//fmt.Printf("%v\n", <-ch) // 从通道获取会报错
}

// 示例3
// 定义Notifier接口
type Notifier interface {
	SendInt(ch chan<- int)
}

// 示例4
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 示例5
type GetIntChan func() <-chan int
